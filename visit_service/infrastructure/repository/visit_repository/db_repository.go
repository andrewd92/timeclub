package visit_repository

import (
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/dao/visit_dao"
	log "github.com/sirupsen/logrus"
	"time"
)

type VisitDbRepository struct {
	dao visit_dao.VisitDao
}

func (r VisitDbRepository) Save(_ *visit.Visit) (*visit.Visit, error) {
	panic("implement me")
}

func (r VisitDbRepository) GetAll() ([]*visit.Visit, error) {
	models, err := r.dao.GetAll()

	if err != nil {
		return nil, err
	}

	cards := make([]*visit.Visit, len(models), len(models))

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

func modelToEntity(model *visit_dao.VisitModel) (*visit.Visit, error) {

	start, err := time.Parse("2006-01-02 15:04:05", model.Start)
	if err != nil {
		return nil, err
	}
	return visit.NewVisit(
		&start,
		model.ClubId,
		order_details.NewOrderDetails([]*event.Event{}),
		model.Comment,
		model.CardId,
		model.ClientName,
	).WithId(model.Id), nil
}
