package domain

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestTimePeriod_CommonSecondsWhenFirstPeriodBeforeSecond(t1 *testing.T) {
	now := time.Now()
	oneSecondDuration := durationSeconds(1)
	periodAEnd := now.Add(oneSecondDuration)
	periodBStart := periodAEnd.Add(oneSecondDuration)
	periodBEnd := periodBStart.Add(oneSecondDuration)

	periodA := NewTimePeriod(now, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 0

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenFirstPeriodAfterSecond(t1 *testing.T) {
	now := time.Now()
	oneSecondDuration := durationSeconds(1)
	periodBEnd := now.Add(oneSecondDuration)
	periodAStart := periodBEnd.Add(oneSecondDuration)
	periodAEnd := periodAStart.Add(oneSecondDuration)

	periodA := NewTimePeriod(periodAStart, periodAEnd)
	periodB := NewTimePeriod(now, periodBEnd)

	var expected int64 = 0

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenSecondPeriodInsideFirst(t1 *testing.T) {
	now := time.Now()
	oneSecondDuration := durationSeconds(1)
	periodBStart := now.Add(oneSecondDuration)
	periodBEnd := periodBStart.Add(oneSecondDuration)
	periodAEnd := periodBEnd.Add(oneSecondDuration)

	periodA := NewTimePeriod(now, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 1

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenFirstPeriodStartBeforeSecond(t1 *testing.T) {
	now := time.Now()
	oneSecondDuration := durationSeconds(1)
	periodBStart := now.Add(oneSecondDuration)
	periodAEnd := periodBStart.Add(oneSecondDuration)
	periodBEnd := periodAEnd.Add(oneSecondDuration)

	periodA := NewTimePeriod(now, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 1

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenFirstPeriodStartAfterSecond(t1 *testing.T) {
	now := time.Now()
	oneSecondDuration := durationSeconds(1)
	periodAStart := now.Add(oneSecondDuration)
	periodBStart := now
	periodAEnd := periodAStart.Add(oneSecondDuration)
	periodBEnd := periodAEnd.Add(oneSecondDuration)

	periodA := NewTimePeriod(periodAStart, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 1

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_CommonSecondsWhenFirstPeriodStartAndEndAfterSecond(t1 *testing.T) {
	now := time.Now()
	oneSecondDuration := durationSeconds(1)
	periodAStart := now.Add(oneSecondDuration)
	periodBStart := now
	periodBEnd := periodAStart.Add(oneSecondDuration)
	periodAEnd := periodBEnd.Add(oneSecondDuration)

	periodA := NewTimePeriod(periodAStart, periodAEnd)
	periodB := NewTimePeriod(periodBStart, periodBEnd)

	var expected int64 = 1

	assert.Equal(t1, expected, periodA.CommonSeconds(periodB))
}

func TestTimePeriod_Duration(t *testing.T) {
	now := time.Now()

	expected := 3600

	period := NewTimePeriod(now, now.Add(durationSeconds(expected)))

	assert.Equal(t, expected, period.Duration())
}

func TestTimePeriod_DurationMinutes(t *testing.T) {
	now := time.Now()

	period := NewTimePeriod(now, now.Add(durationSeconds(3600)))

	assert.Equal(t, 60, period.DurationMinutes())
}

func durationSeconds(seconds int) time.Duration {
	duration, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")

	return duration
}
