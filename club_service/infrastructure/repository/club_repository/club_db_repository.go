package club_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/currency"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/club_dao"
)

type ClubDBRepository struct {
	dao *club_dao.ClubDao
}

func (r ClubDBRepository) GetAll() ([]*club.Club, error) {
	clubModels, err := r.dao.GetAll()
	if err != nil {
		return nil, err
	}

	result := make([]*club.Club, 0, len(clubModels))

	for _, clubModel := range clubModels {
		result = append(result, convertModelToEntity(clubModel))
	}

	return result, nil
}

func (r ClubDBRepository) GetById(id int64) (*club.Club, error) {
	clubModel, err := r.dao.GetById(id)
	if err != nil {
		return nil, err
	}

	return convertModelToEntity(clubModel), nil
}

func convertModelToEntity(model *club_dao.ClubModel) *club.Club {
	return club.NewClub(model.Id, model.Name, model.OpenTime, price_list.Empty(), currency.USD())
}
