package club

import (
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	price2 "github.com/andrewd92/timeclub/club_service/domain/price_list/price"
)

func DefaultClub() *Club {
	currency := price2.NewCurrency(1, "US Dollar", "$")
	prices := []*price2.Price{price2.NewPrice(price2.NewPricePeriod(0, 360), 10, currency)}
	return NewClub(
		1,
		"Club A",
		"12:00",
		price_list.NewPriceList(prices),
		currency,
	)
}
