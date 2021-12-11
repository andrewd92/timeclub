package model

type ClubModel struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	OpenTime    string `db:"open_time"`
	PriceListId int64  `db:"price_list_id"`
	CurrencyId  int64  `db:"currency_id"`
}
