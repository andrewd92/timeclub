package visit_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/client"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/event"
	"github.com/andrewd92/timeclub/club_service/domain/order_details"
	"github.com/andrewd92/timeclub/club_service/domain/visit"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/club_repository"
	"sync"
	"time"
)

type VisitInMemoryRepository struct {
	data map[int64]*visit.Visit

	lock *sync.RWMutex

	clubRepository club.Repository
}

var repository visit.Repository

func Instance() visit.Repository {
	if nil != repository {
		return repository
	}

	repository = &VisitInMemoryRepository{
		data: make(map[int64]*visit.Visit),
		lock: &sync.RWMutex{},

		clubRepository: club_repository.Instance(),
	}

	return repository
}

func (r VisitInMemoryRepository) CreateVisit(clubId int64) (*visit.Visit, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	clubModel, clubErr := r.clubRepository.GetById(clubId)

	if nil != clubErr {
		return nil, clubErr
	}

	now := time.Now()

	newVisit := visit.NewVisit(
		int64(len(r.data)+1),
		&now,
		client.DefaultClient(),
		clubModel,
		order_details.NewOrderDetails(make([]*event.Event, 0)),
		"",
	)

	r.data[newVisit.Id()] = newVisit

	return newVisit, nil
}

func (r VisitInMemoryRepository) GetAll() []*visit.Visit {
	r.lock.RLock()
	defer r.lock.RUnlock()

	result := make([]*visit.Visit, 0, len(r.data))

	for _, value := range r.data {
		result = append(result, value)
	}

	return result
}
