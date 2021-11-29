package price_list_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/price_dao"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/price_list_dao"
)

var repository price_list.Repository

func Instance() price_list.Repository {
	if nil == repository {
		con := connection.Instance()
		repository = PriceListDBRepository{
			priceDao: price_dao.NewPriceDao(con),
			listDao:  price_list_dao.NewPriceListDao(con),
		}
	}

	return repository
}
