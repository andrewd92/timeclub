package test_utils

import "github.com/andrewd92/timeclub/club_service/domain"

func DefaultClub() *domain.Club {
	currency := domain.NewCurrency(1, "US Dollar", "$")
	prices := []*domain.Price{domain.NewPrice(domain.NewPricePeriod(0, 360), 10, currency)}
	return domain.NewClub(
		1,
		"Club A",
		"12:00",
		domain.NewPriceList(prices),
		currency,
	)
}
