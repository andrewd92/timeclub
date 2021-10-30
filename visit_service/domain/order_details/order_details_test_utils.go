package order_details

import (
	"github.com/andrewd92/timeclub/visit_service/domain/event"
)

func DefaultOrderDetails() OrderDetails {
	events := []*event.Event{event.DefaultEvent()}

	return NewOrderDetails(events)
}
