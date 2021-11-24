package currency_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/currency"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/currency_dao"
)

type CurrencyDBRepository struct {
	dao *currency_dao.CurrencyDao
}

func (r CurrencyDBRepository) GetAll() ([]*currency.Currency, error) {
	currencyModels, err := r.dao.GetAll()
	if err != nil {
		return nil, err
	}

	result := make([]*currency.Currency, 0, len(currencyModels))

	for _, currencyModel := range currencyModels {
		result = append(result, convertModelToEntity(currencyModel))
	}

	return result, nil
}

func (r CurrencyDBRepository) GetById(id int64) (*currency.Currency, error) {
	currencyModel, err := r.dao.GetById(id)
	if err != nil {
		return nil, err
	}

	return convertModelToEntity(currencyModel), nil
}

func convertModelToEntity(model *currency_dao.CurrencyModel) *currency.Currency {
	return currency.NewCurrency(model.Id, model.Name, model.ShortName)
}
