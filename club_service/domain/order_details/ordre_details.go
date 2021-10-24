package order_details

import (
	"github.com/andrewd92/timeclub/club_service/domain/event"
)

type OrderDetails struct {
	//sale        int64 //todo change to objects, will be used as client discount
	//certificate int64 // will be used as part of payment
	events []*event.Event
}

func NewOrderDetails(events []*event.Event) *OrderDetails {
	return &OrderDetails{events: events}
}

func (o OrderDetails) Events() []*event.Event {
	return o.events
}
