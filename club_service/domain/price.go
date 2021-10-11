package domain

type Price struct {
	period         *PricePeriod
	valuePerMinute int64
	currency       *Currency
}

func NewPrice(period *PricePeriod, value int64, currency *Currency) *Price {
	return &Price{
		period:         period,
		valuePerMinute: value,
		currency:       currency,
	}
}

func (p Price) Period() *PricePeriod {
	return p.period
}

func (p Price) ValuePerMinute() int64 {
	return p.valuePerMinute
}

func (p Price) Currency() *Currency {
	return p.currency
}

func (p Price) calculate(visit Visit) int64 {
	return p.valuePerMinute * p.period.TimeForPay(visit.Duration())
}
