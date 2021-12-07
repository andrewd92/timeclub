package visit_repository

import (
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	visitPkg "github.com/andrewd92/timeclub/visit_service/domain/visit"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/dao/visit_dao"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/dao/visit_dao/mocks"
	"github.com/andrewd92/timeclub/visit_service/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_GetAll(t *testing.T) {
	dao := new(mocks.VisitDao)
	repository := VisitDbRepository{dao: dao}

	now := time.Now()
	model := visitModel(now.UTC().Format(utils.TimeFormat))
	dao.On("GetAll").Return([]visit_dao.VisitModel{model}, nil)

	expected := visitPkg.NewVisit(
		&now,
		model.ClubId,
		order_details.NewOrderDetails([]*event.Event{}),
		model.Comment,
		model.CardId,
		model.ClientName,
	).WithId(model.Id)

	visits, err := repository.GetAll()
	assert.Nil(t, err)
	assert.NotNil(t, visits)
	assert.True(t, len(visits) == 1)
	actual := visits[0]

	assert.Equal(t, expected.Start().Format(utils.TimeFormat), actual.Start().Format(utils.TimeFormat))
	assert.Equal(t, expected.Id(), actual.Id())
	assert.Equal(t, expected.ClubId(), actual.ClubId())
	assert.Equal(t, expected.CardId(), actual.CardId())
	assert.Equal(t, expected.OrderDetails(), actual.OrderDetails())
	assert.Equal(t, expected.Comment(), actual.Comment())
	assert.Equal(t, expected.ClientName(), actual.ClientName())
}

func Test_Save(t *testing.T) {
	now := time.Now()
	visit := visitPkg.NewVisit(
		&now,
		1,
		order_details.DefaultOrderDetails(),
		"",
		1,
		"Andy",
	)
	id := int64(236)

	dao := new(mocks.VisitDao)
	repository := VisitDbRepository{dao: dao}

	dao.On("Insert", visit).Return(id, nil)

	result, err := repository.Save(visit)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.Id())
	assert.Equal(t, visit.Start().Format(utils.TimeFormat), result.Start().Format(utils.TimeFormat))
}

func visitModel(startTime string) visit_dao.VisitModel {
	return visit_dao.VisitModel{
		Id:           1,
		Start:        startTime,
		ClubId:       1,
		OrderDetails: "[]",
		Comment:      "",
		CardId:       1,
		ClientName:   "Guest",
	}
}
