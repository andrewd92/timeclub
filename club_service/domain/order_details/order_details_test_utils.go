package order_details

import (
	"github.com/andrewd92/timeclub/club_service/domain/event"
)

func DefaultOrderDetails() OrderDetails {
	events := []*event.Event{event.DefaultEvent()}

	return *NewOrderDetails(events)
}
