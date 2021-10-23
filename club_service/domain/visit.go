package domain

import (
	"math"
	"time"
)

type Visit struct {
	id           int64
	start        *time.Time
	end          *time.Time
	client       *Client
	club         *Club
	orderDetails OrderDetails
	comment      string
}

func NewVisit(start *time.Time, end *time.Time, client *Client) *Visit {
	return &Visit{start: start, end: end, client: client}
}

func (v Visit) OrderDetails() OrderDetails {
	return v.orderDetails
}

func (v Visit) Comment() string {
	return v.comment
}

func (v Visit) Id() int64 {
	return v.id
}

func (v Visit) Start() *time.Time {
	return v.start
}

func (v Visit) End() *time.Time {
	if nil != v.end {
		return v.end
	}

	now := time.Now()
	return &now
}

func (v Visit) Client() *Client {
	return v.client
}

func (v Visit) Club() *Club {
	return v.club
}

func (v *Visit) Period() *VisitPeriod {
	return NewVisitPeriod(*v.Start(), *v.End())
}

func (v *Visit) Finish() {
	v.end = v.End()
}

func (v *Visit) Duration() int64 {
	end := time.Now()

	if v.end != nil {
		end = *v.end
	}

	durationOfMinutes := end.Sub(*v.start).Minutes()

	return int64(math.Floor(durationOfMinutes))
}

func (v Visit) calculate(priceList PriceList) (float32, error) {
	visitPeriods, splitErr := v.Period().Split(v.club.openTime)

	if nil != splitErr {
		return 0, splitErr
	}

	var result float32 = 0

	for _, period := range visitPeriods {
		result += period.calculatePrice(priceList, v.orderDetails)
	}

	return result, nil
}
