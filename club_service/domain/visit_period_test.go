package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestVisitPeriod_split(t *testing.T) {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2020-02-06 13:00:00")
	endTime := startTime.Add(time.Hour)
	initPeriod := NewVisitPeriod(startTime, endTime)

	split, _ := initPeriod.Split("12:00")
	assert.Equal(t, []*VisitPeriod{initPeriod}, split)
}

func TestVisitPeriod_splitWhenTwoPeriods(t *testing.T) {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2020-02-06 13:00:00")
	endTime := startTime.Add(24 * time.Hour)
	initPeriod := NewVisitPeriod(startTime, endTime)

	firstPeriodEnd := startTime.Add(23 * time.Hour)
	expected := []*VisitPeriod{
		NewVisitPeriod(startTime, firstPeriodEnd),
		NewVisitPeriodFromMinute(firstPeriodEnd, endTime, 23*60),
	}

	split, _ := initPeriod.Split("12:00")
	assertVisitPeriodsEquals(t, expected, split)
}

func TestVisitPeriod_splitWhenThreePeriods(t *testing.T) {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2020-02-06 10:00:00")
	endTime := startTime.Add(30 * time.Hour)
	initPeriod := NewVisitPeriod(startTime, endTime)

	firstPeriodEnd := startTime.Add(2 * time.Hour)
	secondPeriodEnd := firstPeriodEnd.Add(24 * time.Hour)
	expected := []*VisitPeriod{
		NewVisitPeriod(startTime, firstPeriodEnd),
		NewVisitPeriodFromMinute(firstPeriodEnd, secondPeriodEnd, 2*60),
		NewVisitPeriodFromMinute(secondPeriodEnd, endTime, 26*60),
	}

	split, _ := initPeriod.Split("12:00")
	assertVisitPeriodsEquals(t, expected, split)
}

func TestVisitPeriod_splitOneFullPeriod(t *testing.T) {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2020-02-06 12:00:00")
	endTime := startTime.Add(24 * time.Hour)
	initPeriod := NewVisitPeriod(startTime, endTime)

	expected := []*VisitPeriod{
		NewVisitPeriod(startTime, endTime),
	}

	split, _ := initPeriod.Split("12:00")
	assertVisitPeriodsEquals(t, expected, split)
}

func TestVisitPeriod_SplitTimeError(t *testing.T) {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2020-02-06 12:00:00")
	endTime := startTime.Add(24 * time.Hour)
	initPeriod := NewVisitPeriod(startTime, endTime)

	invalidTimeFormat := "12:00:00"
	_, splitErr := initPeriod.Split(invalidTimeFormat)

	assert.NotNil(t, splitErr)
}

func assertVisitPeriodsEquals(t *testing.T, expected []*VisitPeriod, actual []*VisitPeriod) {
	assert.Equal(t, len(expected), len(actual))

	for i, period := range actual {
		assert.True(t, expected[i].start.Equal(period.start))
		assert.True(t, expected[i].end.Equal(period.end))
		assert.Equal(t, expected[i].firstMinute, period.firstMinute)
	}
}
