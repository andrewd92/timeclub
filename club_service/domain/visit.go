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
	return v.end
}

func (v Visit) Client() *Client {
	return v.client
}

func (v Visit) Club() *Club {
	return v.club
}

func (v *Visit) Finish() {
	if nil != v.end {
		return
	}

	now := time.Now()
	v.end = &now
}

func (v *Visit) Duration() int64 {
	end := time.Now()

	if v.end != nil {
		end = *v.end
	}

	durationOfMinutes := end.Sub(*v.start).Minutes()

	return int64(math.Floor(durationOfMinutes))
}
