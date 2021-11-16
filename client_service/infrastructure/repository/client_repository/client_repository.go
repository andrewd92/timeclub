package client_repository

import (
	"errors"
	"fmt"
	"github.com/andrewd92/timeclub/client_service/domain/client"
	"sync"
	"time"
)

type ClientInMemoryRepository struct {
	data map[int64]*client.Client

	lock *sync.RWMutex
}

var repository client.Repository

func Instance() (client.Repository, error) {
	if nil != repository {
		return repository, nil
	}

	birthday, err := time.Parse("2006-01-02", "1992-02-06")

	if err != nil {
		return nil, err
	}

	repository = &ClientInMemoryRepository{
		data: map[int64]*client.Client{
			int64(999): client.NewClient(
				999,
				"Andrew",
				"D",
				777,
				"andrewd92@gmail.com",
				birthday,
				"",
				1,
				1,
				"Warsaw",
				"",
				time.Now(),
				0,
				999,
			),
		},
		lock: &sync.RWMutex{},
	}

	return repository, nil
}

func (r ClientInMemoryRepository) GetById(id int64) (*client.Client, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	clientModel, ok := r.data[id]
	if false == ok {
		return nil, errors.New(fmt.Sprintf("Client not exists in storage. ID: %d", id))
	}

	return clientModel, nil
}
