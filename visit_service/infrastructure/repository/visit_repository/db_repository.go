package visit_repository

import (
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	visitPkg "github.com/andrewd92/timeclub/visit_service/domain/visit"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/dao/visit_dao"
	"github.com/andrewd92/timeclub/visit_service/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

type VisitDbRepository struct {
	dao visit_dao.VisitDao
}

func (r VisitDbRepository) Save(visit *visitPkg.Visit) (*visitPkg.Visit, error) {
	id, err := r.dao.Insert(visit)
	if err != nil {
		return nil, err
	}

	return visit.WithId(id), nil
}

func (r VisitDbRepository) GetAll() ([]*visitPkg.Visit, error) {
	models, err := r.dao.GetAll()

	if err != nil {
		return nil, err
	}

	cards := make([]*visitPkg.Visit, len(models), len(models))

	for i, model := range models {
		entity, err := modelToEntity(&model)

		if err != nil {
			log.WithError(err).Error("Can not parse visit model")
			continue
		}

		cards[i] = entity
	}

	return cards, nil
}

func modelToEntity(model *visit_dao.VisitModel) (*visitPkg.Visit, error) {
	loc, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		log.WithError(err).Error("can not parse time zone")
	}

	start, err := time.Parse(utils.TimeFormat, model.Start)
	if err != nil {
		return nil, err
	}

	startTime := start.In(loc)
	return visitPkg.NewVisit(
		&startTime,
		model.ClubId,
		order_details.NewOrderDetails([]*event.Event{}),
		model.Comment,
		model.CardId,
		model.ClientName,
	).WithId(model.Id), nil
}
