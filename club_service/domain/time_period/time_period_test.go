package time_period

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimePeriod_CommonSecondsWhenFirstPeriodBeforeSecond(t1 *testing.T) {
	now := time.Now()
	periodAEnd := now.Add(time.Second)
	periodBStart := periodAEnd.Add(time.Second)
	periodBEnd := periodBStart.Add(time.Second)

	periodA := NewTimePeriod(now, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 0

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenFirstPeriodAfterSecond(t1 *testing.T) {
	now := time.Now()
	periodBEnd := now.Add(time.Second)
	periodAStart := periodBEnd.Add(time.Second)
	periodAEnd := periodAStart.Add(time.Second)

	periodA := NewTimePeriod(periodAStart, periodAEnd)
	periodB := NewTimePeriod(now, periodBEnd)

	var expected int64 = 0

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenSecondPeriodInsideFirst(t1 *testing.T) {
	now := time.Now()
	periodBStart := now.Add(time.Second)
	periodBEnd := periodBStart.Add(time.Second)
	periodAEnd := periodBEnd.Add(time.Second)

	periodA := NewTimePeriod(now, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 1

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenFirstPeriodStartBeforeSecond(t1 *testing.T) {
	now := time.Now()
	periodBStart := now.Add(time.Second)
	periodAEnd := periodBStart.Add(time.Second)
	periodBEnd := periodAEnd.Add(time.Second)

	periodA := NewTimePeriod(now, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 1

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenFirstPeriodStartAfterSecond(t1 *testing.T) {
	now := time.Now()
	periodAStart := now.Add(time.Second)
	periodBStart := now
	periodAEnd := periodAStart.Add(time.Second)
	periodBEnd := periodAEnd.Add(time.Second)

	periodA := NewTimePeriod(periodAStart, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 1

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenFirstPeriodStartAndEndAfterSecond(t1 *testing.T) {
	now := time.Now()
	periodAStart := now.Add(time.Second)
	periodBStart := now
	periodBEnd := periodAStart.Add(time.Second)
	periodAEnd := periodBEnd.Add(time.Second)

	periodA := NewTimePeriod(periodAStart, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 1

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_Duration(t *testing.T) {
	now := time.Now()

	expected := 3600

	period := NewTimePeriod(now, now.Add(time.Hour))

	assert.Equal(t, expected, period.Duration())
}

func TestTimePeriod_DurationMinutes(t *testing.T) {
	now := time.Now()

	period := NewTimePeriod(now, now.Add(time.Hour))

	assert.Equal(t, 60, period.DurationMinutes())
}
