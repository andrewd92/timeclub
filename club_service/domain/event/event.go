package event

import (
	"github.com/andrewd92/timeclub/club_service/domain/discount"
	"github.com/andrewd92/timeclub/club_service/domain/time_period"
	"github.com/andrewd92/timeclub/club_service/utils"
)

type Event struct {
	name     string
	tag      string
	discount discount.Discount
	period   time_period.TimePeriod
	price    float32
}

func (e Event) CalculateDiscount(period *time_period.TimePeriod, pricePerMinute float32) float32 {
	return e.discount.From(pricePerMinute * float32(e.minutes(period)))
}

func (e Event) minutes(period *time_period.TimePeriod) int {
	commonMinutes := float64(e.period.CommonSeconds(period)) / 60.0
	return utils.FloorFloat64ToInt(commonMinutes)
}

func NewEvent(name string, tag string, discount discount.Discount, period time_period.TimePeriod) *Event {
	return &Event{name: name, tag: tag, discount: discount, period: period}
}

func (e Event) Name() string {
	return e.name
}

func (e Event) Tag() string {
	return e.tag
}

func (e Event) Discount() discount.Discount {
	return e.discount
}

func (e Event) Period() time_period.TimePeriod {
	return e.period
}
