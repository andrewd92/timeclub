package create

type Request struct {
	Name       string  `json:"name"`
	OpenTime   string  `json:"open_time"`
	PriceList  []Price `json:"price_list_id"`
	CurrencyId int64   `json:"currency_id"`
}

type Price struct {
	PricePeriod    PricePeriod `json:"price_period"`
	ValuePerMinute float32     `json:"value_per_minute"`
}

type PricePeriod struct {
	From int `json:"from"`
	To   int `json:"to"`
}
