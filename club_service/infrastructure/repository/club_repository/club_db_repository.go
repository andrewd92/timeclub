package club_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/currency"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/club_dao"
)

type ClubDBRepository struct {
	dao                 club_dao.ClubDao
	priceListRepository price_list.Repository
}

func (r ClubDBRepository) GetAll() ([]*club.Club, error) {
	clubModels, err := r.dao.GetAll()
	if err != nil {
		return nil, err
	}

	result := make([]*club.Club, 0, len(clubModels))

	for _, clubModel := range clubModels {
		priceList, err := r.priceListRepository.GetById(clubModel.PriceListId)
		if err != nil {
			return nil, err
		}

		result = append(result, convertModelToEntity(clubModel, priceList))
	}

	return result, nil
}

func (r ClubDBRepository) GetById(id int64) (*club.Club, error) {
	clubModel, err := r.dao.GetById(id)
	if err != nil {
		return nil, err
	}

	priceList, err := r.priceListRepository.GetById(clubModel.PriceListId)
	if err != nil {
		return nil, err
	}

	return convertModelToEntity(clubModel, priceList), nil
}

func convertModelToEntity(model *club_dao.ClubModel, priceList *price_list.PriceList) *club.Club {
	return club.NewClub(model.Id, model.Name, model.OpenTime, priceList, currency.USD())
}
