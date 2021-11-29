package club_service

import (
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/club_repository"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/price_list_repository"
)

var service ClubService

func Instance() ClubService {
	if nil == service {
		service = &clubServiceImpl{
			clubRepository:      club_repository.Instance(),
			priceListRepository: price_list_repository.Instance(),
		}
	}

	return service
}
