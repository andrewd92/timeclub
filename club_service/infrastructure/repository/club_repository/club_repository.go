package club_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	"sync"
)

type ClubInMemoryRepository struct {
	data map[int64]*club.Club

	lock *sync.RWMutex
}

var repository *ClubInMemoryRepository

func NewClubInMemoryRepository() *ClubInMemoryRepository {
	if nil != repository {
		return repository
	}

	usd := price.USD()
	priceList := price_list.NewPriceList([]*price.Price{
		price.NewPrice(price.NewPricePeriod(0, 360), 10, usd),
	})

	repository = &ClubInMemoryRepository{
		data: map[int64]*club.Club{
			int64(1): club.NewClub(1, "System", "8:00", priceList, usd),
		},
		lock: &sync.RWMutex{},
	}

	return repository
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
