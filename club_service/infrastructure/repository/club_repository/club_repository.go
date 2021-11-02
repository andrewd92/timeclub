package club_repository

import (
	"errors"
	"fmt"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/currency_repository"
	"sync"
)

type ClubInMemoryRepository struct {
	data map[int64]*club.Club

	lock *sync.RWMutex
}

var repository club.Repository

func Instance() (club.Repository, error) {
	if nil != repository {
		return repository, nil
	}

	currency, getCurrencyErr := currency_repository.Instance().GetById(1)
	if getCurrencyErr != nil {
		return nil, getCurrencyErr
	}

	priceList := price_list.NewPriceList([]*price.Price{
		price.NewPrice(price.NewPricePeriod(0, 360), 10),
	})

	repository = &ClubInMemoryRepository{
		data: map[int64]*club.Club{
			int64(1): club.NewClub(1, "System", "8:00", priceList, currency),
		},
		lock: &sync.RWMutex{},
	}

	return repository, nil
}

func (r ClubInMemoryRepository) GetAll() []*club.Club {
	r.lock.RLock()
	defer r.lock.RUnlock()

	result := make([]*club.Club, 0, len(r.data))

	for _, value := range r.data {
		result = append(result, value)
	}

	return result
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
