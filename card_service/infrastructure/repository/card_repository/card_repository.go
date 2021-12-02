package card_repository

import (
	"errors"
	"fmt"
	"github.com/andrewd92/timeclub/card_service/domain/card"
	"sync"
)

type CardInMemoryRepository struct {
	data map[int64]*card.Card

	lock *sync.RWMutex
}

func (r CardInMemoryRepository) GetById(id int64) (*card.Card, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	cardModel, ok := r.data[id]
	fmt.Println(ok)
	if false == ok {
		return nil, errors.New("card not exists")
	}

	return cardModel, nil
}

func (r CardInMemoryRepository) Save(card *card.Card) (*card.Card, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	cardModel := card.WithId(int64(len(r.data) + 1))

	r.data[cardModel.Id()] = cardModel

	return cardModel, nil
}

func (r CardInMemoryRepository) All() ([]*card.Card, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := make([]*card.Card, 0, len(r.data))

	for _, value := range r.data {
		result = append(result, value)
	}

	return result, nil
}
