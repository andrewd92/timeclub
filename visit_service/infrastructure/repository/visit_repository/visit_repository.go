package visit_repository

import (
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"sync"
)

type VisitInMemoryRepository struct {
	data map[int64]*visit.Visit

	lock *sync.RWMutex
}

func (r VisitInMemoryRepository) Save(visit *visit.Visit) (*visit.Visit, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	id := int64(len(r.data) + 1)

	model := visit.WithId(id)

	r.data[model.Id()] = model

	return model, nil
}

func (r VisitInMemoryRepository) GetAll() ([]*visit.Visit, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	result := make([]*visit.Visit, 0, len(r.data))

	for _, value := range r.data {
		result = append(result, value)
	}

	return result, nil
}
