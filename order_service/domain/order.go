package domain

import (
	"time"
)

type Order struct {
	Id     int64
	Status OrderStatus
	Start  time.Time
	End    *time.Time
	Visits []int64
}

func (o *Order) Cancel() {
	if o.Status != Open {
		return
	}

	o.Status = Cancel
}

func (o *Order) Pay() {
	if o.Status != Open {
		return
	}

	o.Status = Paid
}
