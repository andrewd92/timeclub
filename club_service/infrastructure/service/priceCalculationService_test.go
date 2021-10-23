package service

import (
	"github.com/andrewd92/timeclub/club_service/domain"
	"github.com/andrewd92/timeclub/club_service/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPriceCalculationService_calculate(t *testing.T) {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2020-02-06 13:00:00")
	endTime := startTime.Add(time.Hour)
	visit := domain.NewVisit(&startTime, &endTime, test_utils.DefaultClient())
	club := test_utils.DefaultClub()

	service := NewPriceCalculationService()

	price, err := service.calculate(visit, club)
	expected := float32(600)

	assert.Nil(t, err)
	assert.Equal(t, expected, price)
}
