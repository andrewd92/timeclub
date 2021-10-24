package service

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/visit"
	"time"
)

type PriceCalculationService struct {
}

func NewPriceCalculationService() *PriceCalculationService {
	return &PriceCalculationService{}
}

func (s PriceCalculationService) calculate(visit *visit.Visit, club *club.Club) (float32, error) {
	now := time.Now()
	visitPeriod := visit.Period(now)
	visitPeriods, splitErr := visitPeriod.Split(club.OpenTime())
	if splitErr != nil {
		return 0, splitErr
	}

	var price float32 = 0

	for _, period := range visitPeriods {
		price += club.PriceList().Calculate(period.Duration())
	}

	return price, nil
}
