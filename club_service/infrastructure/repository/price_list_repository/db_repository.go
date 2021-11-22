package price_list_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/price_dao"
)

type PriceListDBRepository struct {
	dao *price_dao.PriceDao
}

func (r PriceListDBRepository) GetById(priceListId int64) (*price_list.PriceList, error) {
	priceModels, err := r.dao.GetByPriceListId(priceListId)
	if err != nil {
		return nil, err
	}

	result := make([]*price.Price, 0, len(priceModels))

	for _, priceModel := range priceModels {
		result = append(result, convertModelToEntity(priceModel))
	}

	return price_list.NewPriceList(result), nil
}

func convertModelToEntity(model *price_dao.PriceModel) *price.Price {
	return price.NewPrice(
		price.NewPricePeriod(model.PeriodFrom, model.PeriodTo),
		model.ValuePerMinute,
	)
}
