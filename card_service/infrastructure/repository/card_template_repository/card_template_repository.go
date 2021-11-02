package card_template_repository

import (
	"errors"
	"fmt"
	"github.com/andrewd92/timeclub/card_service/domain/card/card_template"
	"sync"
)

type CardTemplateInMemoryRepository struct {
	data map[int64]*card_template.CardTemplate

	lock *sync.RWMutex
}

var repository card_template.Repository

func Instance() card_template.Repository {
	if nil != repository {
		return repository
	}

	repository = &CardTemplateInMemoryRepository{
		data: make(map[int64]*card_template.CardTemplate),
		lock: &sync.RWMutex{},
	}

	return repository
}

func (r CardTemplateInMemoryRepository) GetById(id int64) (*card_template.CardTemplate, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	cardTemplateModel, ok := r.data[id]
	fmt.Println(ok)
	if false == ok {
		return nil, errors.New("card not exists")
	}

	return cardTemplateModel, nil
}

func (r CardTemplateInMemoryRepository) Save(cardTemplate *card_template.CardTemplate) (*card_template.CardTemplate, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	cardTemplateModel := cardTemplate.WithId(int64(len(r.data) + 1))

	r.data[cardTemplateModel.Id()] = cardTemplateModel

	return cardTemplateModel, nil
}

func (r CardTemplateInMemoryRepository) GetAll() []*card_template.CardTemplate {
	r.lock.RLock()
	defer r.lock.RUnlock()

	result := make([]*card_template.CardTemplate, 0, len(r.data))

	for _, value := range r.data {
		result = append(result, value)
	}

	return result
}
