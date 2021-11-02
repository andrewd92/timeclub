package price

type priceJson struct {
	PricePeriod    map[string]int `json:"price_period"`
	ValuePerMinute float32        `json:"value_per_minute"`
	Currency       string         `json:"currency"`
}

func (p Price) Marshal() interface{} {
	pricePeriod := map[string]int{
		"from": p.pricePeriod.from,
		"to":   p.pricePeriod.to,
	}

	return priceJson{
		PricePeriod:    pricePeriod,
		ValuePerMinute: p.valuePerMinute,
	}
}
