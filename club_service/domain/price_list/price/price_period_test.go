package price

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPricePeriod_TimeForPay(t *testing.T) {
	period := NewPricePeriod(0, 60)

	durationOfVisit := 120

	assert.Equal(t, period.totalTime(), period.TimeForPay(durationOfVisit))
}

func TestPricePeriod_TimeForPayWhenDurationLessThenEndOfPeriod(t *testing.T) {
	period := NewPricePeriod(0, 60)

	durationOfVisit := 30

	assert.Equal(t, durationOfVisit, period.TimeForPay(durationOfVisit))
}

func TestPricePeriod_TimeForPayShouldReturnZeroWhenDurationLessThenStartOfPeriod(t *testing.T) {
	period := NewPricePeriod(120, 180)

	durationOfVisit := 30

	assert.Equal(t, 0, period.TimeForPay(durationOfVisit))
}

func TestPricePeriod_totalTime(t *testing.T) {
	period := NewPricePeriod(0, 60)

	assert.Equal(t, 60, period.totalTime())
}
