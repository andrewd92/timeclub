package event

import (
	discountPkg "github.com/andrewd92/timeclub/visit_service/domain/discount"
	"github.com/andrewd92/timeclub/visit_service/domain/time_period"
	"time"
)

const DefaultDiscountPercent = 10
const DefaultDiscountPerMinute = 1

func DefaultEvent() *Event {
	now := time.Now()

	return DefaultEventFrom(&now)
}

func DefaultEventFrom(start *time.Time) *Event {
	discount := discountPkg.NewDiscount(float32(DefaultDiscountPercent))

	return NewEvent("A", "#a", *discount, *time_period.NewTimePeriod(*start, start.Add(24*time.Hour)))
}
