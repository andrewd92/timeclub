package visit

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/event"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	"github.com/andrewd92/timeclub/club_service/domain/visit/visit_period"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestVisit_Duration(t *testing.T) {
	now := time.Now()
	start := now.Add(-2 * time.Hour)
	visit := DefaultVisitFrom(&start)

	assert.Equal(t, 120, visit.Duration(now))
}

func TestVisit_Period(t *testing.T) {
	now := time.Now()
	start := now.Add(-2 * time.Hour)
	visit := DefaultVisitFrom(&start)

	expected := visit_period.NewVisitPeriod(start, now)
	actual := visit.Period(now)

	assert.Equal(t, expected.Start(), actual.Start())
	assert.Equal(t, expected.End(), actual.End())
	assert.Equal(t, expected.FirstMinute(), actual.FirstMinute())
}

func TestVisit_CalculatePrice(t *testing.T) {
	visit := DefaultVisit()

	end := visit.Start().Add(1 * time.Hour)

	discountCoefficient := float32(0.9)

	expected := float32(price.DefaultPriceValue*60-event.DefaultDiscountPerMinute*60) * discountCoefficient

	actual, err := visit.CalculatePrice(price_list.DefaultPriceList(), end)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestVisit_CalculatePriceShouldReturnErrorWhenWrongClubOpenTime(t *testing.T) {
	visit := VisitWithClub(club.ClubWithWrongOpenTime())

	end := visit.Start().Add(1 * time.Hour)

	_, err := visit.CalculatePrice(price_list.DefaultPriceList(), end)

	assert.NotNil(t, err)
}
