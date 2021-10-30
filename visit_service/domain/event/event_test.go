package event

import (
	"github.com/andrewd92/timeclub/visit_service/domain/time_period"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEvent_minutes(t *testing.T) {
	event := DefaultEvent()

	timePeriodStart := event.Period().Start().Add(time.Hour)
	timePeriod := time_period.NewTimePeriod(timePeriodStart, timePeriodStart.Add(time.Hour))

	assert.Equal(t, 60, event.minutes(timePeriod))
}

func TestEvent_CalculateDiscount(t *testing.T) {
	event := DefaultEvent()

	timePeriodStart := event.Period().Start().Add(time.Hour)
	timePeriod := time_period.NewTimePeriod(timePeriodStart, timePeriodStart.Add(time.Hour))

	expected := float32(60)

	assert.Equal(t, expected, event.CalculateDiscount(timePeriod, 10))
}
