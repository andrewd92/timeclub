package visit_period

import (
	"github.com/andrewd92/timeclub/club_service/api"
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	"github.com/andrewd92/timeclub/visit_service/domain/time_period"
	"github.com/andrewd92/timeclub/visit_service/utils"
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

func TestVisitPeriod_Duration(t *testing.T) {
	start := time.Now()
	period := NewVisitPeriod(start, start.Add(5*time.Minute))

	assert.Equal(t, 5, period.Duration())
}

func TestVisitPeriod_Period(t *testing.T) {
	start := time.Now()
	period := NewVisitPeriod(start, start.Add(1*time.Hour))

	expected := time_period.NewTimePeriod(
		start.Add(time.Minute),
		start.Add(2*time.Minute),
	)

	actual := period.timePeriod(&api.PricePeriod{From: 1, To: 2})

	assert.True(t, expected.Start().Equal(actual.Start()))
	assert.True(t, expected.End().Equal(actual.End()))
}

func TestVisitPeriod_PeriodWhenPricePeriodGreaterThenVisitPeriod(t *testing.T) {
	start := time.Now()
	period := NewVisitPeriod(start, start.Add(time.Hour))

	expected := time_period.NewTimePeriod(
		start,
		start.Add(time.Hour),
	)

	actual := period.timePeriod(&api.PricePeriod{From: 0, To: 60 * 24})

	assert.True(t, expected.Start().Equal(actual.Start()))
	assert.True(t, expected.End().Equal(actual.End()))
}

func TestVisitPeriod_CalculatePrice(t *testing.T) {
	now := time.Now()
	hoursBeforeStart := int64(8)
	visitStart := now.Add(time.Duration(-1*hoursBeforeStart) * time.Hour)
	period := NewVisitPeriod(visitStart, now)
	prices := []*api.Price{utils.DefaultPrice()}

	eventStart := visitStart.Add(30 * time.Minute)
	events := []*event.Event{event.DefaultEventFrom(&eventStart)}
	details := order_details.NewOrderDetails(events)

	actual := period.CalculatePrice(prices, details)

	discountFromEvent := float32(utils.DefaultPricePeriodDurationMinutes) - 30
	expected := utils.DefaultPriceValue*float32(utils.DefaultPricePeriodDurationMinutes) - discountFromEvent

	assert.Equal(t, expected, actual)
}

func TestCalculatePriceForPeriod(t *testing.T) {
	price := utils.DefaultPrice()

	now := time.Now()

	timePeriod := time_period.NewTimePeriod(now, now.Add(time.Hour))

	assert.Equal(t, float32(600), calculatePriceForPeriod(price, timePeriod))
}

func assertVisitPeriodsEquals(t *testing.T, expected []*VisitPeriod, actual []*VisitPeriod) {
	assert.Equal(t, len(expected), len(actual))

	for i, period := range actual {
		assert.True(t, expected[i].start.Equal(period.start))
		assert.True(t, expected[i].end.Equal(period.end))
		assert.Equal(t, expected[i].firstMinute, period.firstMinute)
	}
}
