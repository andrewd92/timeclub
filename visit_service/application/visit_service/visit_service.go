package visit_service

import (
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	"github.com/andrewd92/timeclub/visit_service/domain/price_list"
	"github.com/andrewd92/timeclub/visit_service/domain/price_list/price"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"time"
)

type VisitService interface {
	Create(clubId int64, cardId int64) (interface{}, error)
}

type visitServiceImpl struct {
	visitRepository visit.Repository
}

func (v visitServiceImpl) Create(clubId int64, cardId int64) (interface{}, error) {
	now := time.Now()

	newVisit := visit.NewVisit(
		&now,
		clubId,
		order_details.NewOrderDetails(make([]*event.Event, 0)),
		"",
		cardId,
		"Guest",
	)

	visitModel, saveErr := v.visitRepository.Save(newVisit)

	if nil != saveErr {
		return nil, saveErr
	}

	return visitModel.Marshal(now, price_list.DefaultPriceList(), price.USD()), nil
}
