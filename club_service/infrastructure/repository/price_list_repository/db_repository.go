package price_list_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/price_dao"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/price_list_dao"
	"time"
)

type PriceListDBRepository struct {
	priceDao *price_dao.PriceDao
	listDao  *price_list_dao.PriceListDao
}

func (r PriceListDBRepository) GetById(priceListId int64) (*price_list.PriceList, error) {
	priceModels, err := r.priceDao.GetByPriceListId(priceListId)
	if err != nil {
		return nil, err
	}

	result := make([]*price.Price, 0, len(priceModels))

	for _, priceModel := range priceModels {
		result = append(result, convertModelToEntity(priceModel))
	}

	return price_list.NewPriceListWithId(priceListId, result), nil
}

func (r PriceListDBRepository) Save(list *price_list.PriceList) (*price_list.PriceList, error) {
	if list.Id() != 0 {
		err := r.priceDao.DeletePriceList(list.Id())
		if err != nil {
			return nil, err
		}
	} else {
		id, err := r.listDao.Insert(time.Now().String())
		if err != nil {
			return nil, err
		}
		list = list.WithId(id)
	}

	for _, priceEntity := range list.Prices() {
		_, err := r.priceDao.Insert(priceEntity, list.Id())
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func convertModelToEntity(model *price_dao.PriceModel) *price.Price {
	return price.NewPrice(
		price.NewPricePeriod(model.PeriodFrom, model.PeriodTo),
		model.ValuePerMinute,
	)
}
