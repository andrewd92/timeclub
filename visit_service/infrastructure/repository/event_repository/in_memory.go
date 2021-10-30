package event_repository

import (
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"sync"
)

type EventInMemoryRepository struct {
	data map[int64]*event.Event

	lock *sync.RWMutex
}

var repository event.Repository

func Instance() event.Repository {
	if nil != repository {
		return repository
	}

	repository = &EventInMemoryRepository{
		data: make(map[int64]*event.Event),
		lock: &sync.RWMutex{},
	}

	return repository
}

func (r EventInMemoryRepository) Save(event *event.Event) (*event.Event, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	eventModel := event.WithId(int64(len(r.data) + 1))

	r.data[eventModel.Id()] = eventModel

	return eventModel, nil
}

func (r EventInMemoryRepository) GetAll() []*event.Event {
	r.lock.RLock()
	defer r.lock.RUnlock()

	result := make([]*event.Event, 0, len(r.data))

	for _, value := range r.data {
		result = append(result, value)
	}

	return result
}
