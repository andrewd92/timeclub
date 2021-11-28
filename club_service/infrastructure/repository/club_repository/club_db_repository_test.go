package club_repository

import (
	"errors"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	priceListMock "github.com/andrewd92/timeclub/club_service/domain/price_list/mocks"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/club_dao"
	daoMock "github.com/andrewd92/timeclub/club_service/infrastructure/dao/club_dao/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const defaultPriceListId int64 = 1

func Test_GetById(t *testing.T) {
	var id int64 = 8
	model := createModel(id)
	priceList := price_list.DefaultPriceList()

	dao := new(daoMock.ClubDao)
	dao.On("GetById", id).Return(model, nil)
	priceListRepository := new(priceListMock.Repository)
	priceListRepository.On("GetById", defaultPriceListId).Return(priceList, nil)

	repo := ClubDBRepository{dao: dao, priceListRepository: priceListRepository}

	entity, err := repo.GetById(id)

	expected := convertModelToEntity(model, priceList)

	assert.Nil(t, err)
	assert.Equal(t, expected, entity)
}

func Test_GetById_WhenDbError(t *testing.T) {
	var id int64 = 61
	msg := "test error"

	dao := new(daoMock.ClubDao)
	dao.On("GetById", id).Return(nil, errors.New(msg))
	priceListRepository := new(priceListMock.Repository)

	repo := ClubDBRepository{dao: dao, priceListRepository: priceListRepository}

	entity, err := repo.GetById(id)

	assert.Nil(t, entity)
	assert.Equal(t, msg, err.Error())
}

func Test_GetById_WhenPriceListError(t *testing.T) {
	var id int64 = 23
	model := createModel(id)
	msg := "test price list error"

	dao := new(daoMock.ClubDao)
	dao.On("GetById", id).Return(model, nil)
	priceListRepository := new(priceListMock.Repository)
	priceListRepository.On("GetById", defaultPriceListId).Return(nil, errors.New(msg))

	repo := ClubDBRepository{dao: dao, priceListRepository: priceListRepository}

	entity, err := repo.GetById(id)

	assert.Nil(t, entity)
	assert.Equal(t, msg, err.Error())
}

func Test_GetAll(t *testing.T) {
	var id int64 = 31
	model := createModel(id)
	priceList := price_list.DefaultPriceList()

	dao := new(daoMock.ClubDao)
	dao.On("GetAll").Return([]*club_dao.ClubModel{model}, nil)
	priceListRepository := new(priceListMock.Repository)
	priceListRepository.On("GetById", defaultPriceListId).Return(priceList, nil)

	repo := ClubDBRepository{dao: dao, priceListRepository: priceListRepository}

	clubs, err := repo.GetAll()

	expected := []*club.Club{convertModelToEntity(model, priceList)}

	assert.Nil(t, err)
	assert.Equal(t, expected, clubs)
}

func Test_GetAll_WhenDbError(t *testing.T) {
	msg := "test get all error"

	dao := new(daoMock.ClubDao)
	dao.On("GetAll").Return(nil, errors.New(msg))
	priceListRepository := new(priceListMock.Repository)

	repo := ClubDBRepository{dao: dao, priceListRepository: priceListRepository}

	entity, err := repo.GetAll()

	assert.Nil(t, entity)
	assert.Equal(t, msg, err.Error())
}

func Test_GetAll_WhenPriceListError(t *testing.T) {
	var id int64 = 23
	model := createModel(id)
	msg := "test price list error"

	dao := new(daoMock.ClubDao)
	dao.On("GetAll").Return([]*club_dao.ClubModel{model}, nil)
	priceListRepository := new(priceListMock.Repository)
	priceListRepository.On("GetById", defaultPriceListId).Return(nil, errors.New(msg))

	repo := ClubDBRepository{dao: dao, priceListRepository: priceListRepository}

	entity, err := repo.GetAll()

	assert.Nil(t, entity)
	assert.Equal(t, msg, err.Error())
}

func createModel(id int64) *club_dao.ClubModel {
	return &club_dao.ClubModel{
		Id:          id,
		Name:        "t",
		OpenTime:    "12:00",
		PriceListId: defaultPriceListId,
		CurrencyId:  1,
	}
}
