package order_details

import (
	"github.com/andrewd92/timeclub/club_service/domain/event"
)

type OrderDetails struct {
	events []*event.Event
}

func NewOrderDetails(events []*event.Event) *OrderDetails {
	return &OrderDetails{events: events}
}

func (o OrderDetails) Events() []*event.Event {
	return o.events
}
