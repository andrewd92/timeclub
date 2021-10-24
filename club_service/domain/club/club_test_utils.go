package club

import (
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
)

func DefaultClub() *Club {
	return NewClub(
		1,
		"Club A",
		"12:00",
		price_list.DefaultPriceList(),
		price.USD(),
	)
}

func ClubWithWrongOpenTime() *Club {
	return NewClub(
		1,
		"Club A",
		"WrongOpenTime",
		price_list.DefaultPriceList(),
		price.USD(),
	)
}
