package event

import (
	"github.com/andrewd92/timeclub/club_service/domain/discount"
	"github.com/andrewd92/timeclub/club_service/domain/time_period"
	"time"
)

const DefaultDiscountPercent = 10
const DefaultDiscountPerMinute = 1

func DefaultEvent(start *time.Time) *Event {
	now := time.Now()

	return DefaultEventFrom(&now)
}

func DefaultEventFrom(start *time.Time) *Event {
	discount := discount.NewDiscount(float32(DefaultDiscountPercent))

	return NewEvent("A", "#a", *discount, *time_period.NewTimePeriod(*start, start.Add(24*time.Hour)))
}
