package club

import (
	"github.com/andrewd92/timeclub/club_service/domain/currency"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
)

func DefaultClub() *Club {
	return NewClubWithId(
		1,
		"Club A",
		"12:00",
		price_list.DefaultPriceList(),
		currency.USD(),
	)
}
