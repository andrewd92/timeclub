package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMaxTimeShouldReturnMaxTime(t *testing.T) {
	now := time.Now()
	oneSecondDuration, _ := time.ParseDuration("1s")
	after := now.Add(oneSecondDuration)

	assert.Equal(t, after, MaxTime(now, after))
	assert.Equal(t, after, MaxTime(after, now))
}

func TestMinTimeShouldReturnMixTime(t *testing.T) {
	now := time.Now()
	oneSecondDuration, _ := time.ParseDuration("1s")
	after := now.Add(oneSecondDuration)

	assert.Equal(t, now, MinTime(now, after))
	assert.Equal(t, now, MinTime(after, now))
}

func TestSplitTimeString(t *testing.T) {
	expectedHour := 12
	expectedMinute := 34
	timeString := "12:34"

	actualHour, actualMinute, actualErr := SplitTimeString(timeString)
	assert.Equal(t, expectedHour, actualHour)
	assert.Equal(t, expectedMinute, actualMinute)
	assert.Nil(t, actualErr)
}

func TestSplitTimeStringWhenFormatError(t *testing.T) {
	timeString := "12-34"

	actualHour, actualMinute, actualErr := SplitTimeString(timeString)
	assert.Equal(t, 0, actualHour)
	assert.Equal(t, 0, actualMinute)
	assert.NotNil(t, actualErr)
}

func TestSplitTimeStringWhenParseHourError(t *testing.T) {
	timeString := "aa-34"

	actualHour, actualMinute, actualErr := SplitTimeString(timeString)
	assert.Equal(t, 0, actualHour)
	assert.Equal(t, 0, actualMinute)
	assert.NotNil(t, actualErr)
}

func TestSplitTimeStringWhenParseMinuteError(t *testing.T) {
	timeString := "12-aa"

	actualHour, actualMinute, actualErr := SplitTimeString(timeString)
	assert.Equal(t, 0, actualHour)
	assert.Equal(t, 0, actualMinute)
	assert.NotNil(t, actualErr)
}
