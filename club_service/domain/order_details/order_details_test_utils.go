package order_details

import (
	"github.com/andrewd92/timeclub/club_service/domain/event"
	"time"
)

func DefaultOrderDetails() OrderDetails {
	now := time.Now()
	events := []*event.Event{event.DefaultEvent(&now)}

	return *NewOrderDetails(events)
}
