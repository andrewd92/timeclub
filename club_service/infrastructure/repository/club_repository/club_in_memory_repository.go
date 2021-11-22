package club_repository

import (
	"errors"
	"fmt"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"sync"
)

type ClubInMemoryRepository struct {
	data map[int64]*club.Club

	lock *sync.RWMutex
}

func (r ClubInMemoryRepository) GetAll() ([]*club.Club, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	result := make([]*club.Club, 0, len(r.data))

	for _, value := range r.data {
		result = append(result, value)
	}

	return result, nil
}

func (r ClubInMemoryRepository) GetById(id int64) (*club.Club, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	clubModel, ok := r.data[id]
	if false == ok {
		return nil, errors.New(fmt.Sprintf("Club not exists in storage. ID: %d", id))
	}

	return clubModel, nil
}
