package price

import (
	"github.com/andrewd92/timeclub/club_service/domain/time_period"
)

type Price struct {
	pricePeriod    *PricePeriod
	valuePerMinute float32
	currency       *Currency
}

func NewPrice(period *PricePeriod, value float32, currency *Currency) *Price {
	return &Price{
		pricePeriod:    period,
		valuePerMinute: value,
		currency:       currency,
	}
}

func (p Price) PricePeriod() *PricePeriod {
	return p.pricePeriod
}

func (p Price) ValuePerMinute() float32 {
	return p.valuePerMinute
}

func (p Price) Currency() *Currency {
	return p.currency
}

func (p Price) Calculate(durationMinutes int) float32 {
	return p.valuePerMinute * float32(p.pricePeriod.TimeForPay(durationMinutes))
}

func (p Price) CalculateForPeriod(period time_period.TimePeriod) float32 {
	return float32(period.DurationMinutes()) * p.valuePerMinute
}

func (p Price) Max() float32 {
	return p.valuePerMinute * float32(p.pricePeriod.totalTime())
}
