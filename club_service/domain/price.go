package domain

type Price struct {
	period         *PricePeriod
	valuePerMinute float32
	currency       *Currency
}

func NewPrice(period *PricePeriod, value float32, currency *Currency) *Price {
	return &Price{
		period:         period,
		valuePerMinute: value,
		currency:       currency,
	}
}

func (p Price) Period() *PricePeriod {
	return p.period
}

func (p Price) ValuePerMinute() float32 {
	return p.valuePerMinute
}

func (p Price) Currency() *Currency {
	return p.currency
}

func (p Price) Calculate(visitPeriod *VisitPeriod) float32 {
	return p.valuePerMinute * float32(p.period.TimeForPay(visitPeriod.Duration()))
}

func (p Price) CalculateForPeriod(period TimePeriod) float32 {
	return float32(period.DurationMinutes()) * p.valuePerMinute
}

func (p Price) Max() float32 {
	return p.valuePerMinute * float32(p.period.totalTime())
}
