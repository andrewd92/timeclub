package visit

import (
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	"time"
)

type Factory interface {
	Create(clubId int64, cardId int64) *Visit
}

type FactoryImpl struct {
}

func (f FactoryImpl) Create(clubId int64, cardId int64) *Visit {
	now := time.Now()

	return NewVisit(
		&now,
		clubId,
		order_details.NewOrderDetails(make([]*event.Event, 0)),
		"",
		cardId,
		"Guest",
	)
}
