package club_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/club_dao"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/currency_repository"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/price_list_repository"
	"sync"
)

var repository club.Repository

func Instance() (club.Repository, error) {
	if nil != repository {
		return repository, nil
	}

	repository = &ClubDBRepository{
		dao:                 &club_dao.ClubDao{},
		priceListRepository: price_list_repository.Instance(),
	}

	return repository, nil
}

func InMemoryInstance() (club.Repository, error) {
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
