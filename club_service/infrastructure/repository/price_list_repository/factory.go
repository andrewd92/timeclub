package price_list_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/price_dao"
)

var repository price_list.Repository

func Instance() price_list.Repository {
	if nil == repository {
		repository = PriceListDBRepository{dao: &price_dao.PriceDao{}}
	}

	return repository
}
