package domain

type Price struct {
	period   *PricePeriod
	value    int64
	currency *Currency
}

func NewPrice(period *PricePeriod, value int64, currency *Currency) *Price {
	return &Price{
		period:   period,
		value:    value,
		currency: currency,
	}
}
