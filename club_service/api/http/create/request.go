package create

type Request struct {
	Name        string `json:"name"`
	OpenTime    string `json:"open_time"`
	PriceListId int64  `json:"price_list_id"`
	CurrencyId  int64  `json:"currency_id"`
}
