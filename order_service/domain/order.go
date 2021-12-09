package domain

import (
	"github.com/andrewd92/timeclub/order_service/domain/order_status"
	"time"
)

type Order struct {
	Id     int64
	Status order_status.OrderStatus
	Start  time.Time
	End    *time.Time
	Visits []int64
}

func (o *Order) Cancel() {
	if o.Status != order_status.Open {
		return
	}

	o.Status = order_status.Cancel
}

func (o *Order) Pay() {
	if o.Status != order_status.Open {
		return
	}

	o.Status = order_status.Paid
}
