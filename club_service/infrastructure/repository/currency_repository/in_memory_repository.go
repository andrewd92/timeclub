package currency_repository

import (
	"errors"
	"fmt"
	currencyPkg "github.com/andrewd92/timeclub/club_service/domain/currency"
	"sync"
)

type CurrencyInMemoryRepository struct {
	data map[int64]*currencyPkg.Currency

	lock *sync.RWMutex
}

func (r CurrencyInMemoryRepository) GetAll() ([]*currencyPkg.Currency, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	result := make([]*currencyPkg.Currency, 0, len(r.data))

	for _, value := range r.data {
		result = append(result, value)
	}

	return result, nil
}

func (r CurrencyInMemoryRepository) GetById(id int64) (*currencyPkg.Currency, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	currency, ok := r.data[id]
	if false == ok {
		return nil, errors.New(fmt.Sprintf("Currency not exists in storage. ID: %d", id))
	}

	return currency, nil
}
