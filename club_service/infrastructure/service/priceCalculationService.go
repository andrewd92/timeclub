package service

import "github.com/andrewd92/timeclub/club_service/domain"

type PriceCalculationService struct {
}

func NewPriceCalculationService() *PriceCalculationService {
	return &PriceCalculationService{}
}

func (s PriceCalculationService) calculate(visit *domain.Visit, club *domain.Club) (float32, error) {
	visitPeriod := visit.Period()
	visitPeriods, splitErr := visitPeriod.Split(club.OpenTime())
	if splitErr != nil {
		return 0, splitErr
	}

	var price float32 = 0

	for _, period := range visitPeriods {
		price += club.PriceList().Calculate(period)
	}

	return price, nil
}
