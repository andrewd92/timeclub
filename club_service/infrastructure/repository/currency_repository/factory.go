package currency_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/currency"
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/currency_dao"
)

func Instance() currency.Repository {
	return CurrencyDBRepository{dao: currency_dao.NewCurrencyDao(connection.Instance())}
}
