package domain

import (
	"github.com/andrewd92/timeclub/club_service/utils"
)

type Event struct {
	name     string
	tag      string
	discount Discount
	period   TimePeriod
	price    float32
}

func (e Event) Name() string {
	return e.name
}

func (e Event) Tag() string {
	return e.tag
}

func (e Event) Discount() Discount {
	return e.discount
}

func (e Event) Period() TimePeriod {
	return e.period
}

func (e Event) calculateDiscount(period *TimePeriod, pricePerMinute float32) float32 {
	return pricePerMinute * e.discount.multiplier() * float32(e.minutes(period))
}

func (e Event) minutes(period *TimePeriod) int {
	commonMinutes := float64(e.period.CommonSeconds(period)) / 60.0
	return utils.FloorFloat64ToInt(commonMinutes)
}

func NewEvent(name string, tag string, discount Discount, period TimePeriod) *Event {
	return &Event{name: name, tag: tag, discount: discount, period: period}
}
