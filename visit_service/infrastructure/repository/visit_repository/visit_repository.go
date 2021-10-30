package visit_repository

import (
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"strconv"
	"sync"
	"time"
)

type VisitInMemoryRepository struct {
	data map[int64]*visit.Visit

	lock *sync.RWMutex
}

var repository visit.Repository

func Instance() visit.Repository {
	if nil != repository {
		return repository
	}

	repository = &VisitInMemoryRepository{
		data: make(map[int64]*visit.Visit),
		lock: &sync.RWMutex{},
	}

	return repository
}

func (r VisitInMemoryRepository) CreateVisit(clubId int64, cardId int64) (*visit.Visit, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	now := time.Now()

	id := int64(len(r.data) + 1)
	newVisit := visit.NewVisit(
		id,
		&now,
		clubId,
		order_details.NewOrderDetails(make([]*event.Event, 0)),
		"",
		cardId,
		"Guest"+strconv.FormatInt(id, 10),
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
