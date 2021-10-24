package price

import (
	"github.com/andrewd92/timeclub/club_service/domain/time_period"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPrice_CalculateForPeriod(t *testing.T) {
	price := DefaultPrice()

	now := time.Now()

	timePeriod := time_period.NewTimePeriod(now, now.Add(time.Hour))
	price.CalculateForPeriod(*timePeriod)

	assert.Equal(t, float32(600), price.CalculateForPeriod(*timePeriod))
}

func TestPrice_Max(t *testing.T) {
	price := DefaultPrice()

	assert.Equal(t, float32(600), price.Max())
}
