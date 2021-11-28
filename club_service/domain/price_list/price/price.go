package price

type Price struct {
	pricePeriod    *PricePeriod
	valuePerMinute float32
}

func NewPrice(period *PricePeriod, value float32) *Price {
	return &Price{
		pricePeriod:    period,
		valuePerMinute: value,
	}
}

func (p Price) PricePeriod() *PricePeriod {
	return p.pricePeriod
}

func (p Price) ValuePerMinute() float32 {
	return p.valuePerMinute
}
