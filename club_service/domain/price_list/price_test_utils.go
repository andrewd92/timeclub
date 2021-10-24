package price_list

import (
	price2 "github.com/andrewd92/timeclub/club_service/domain/price_list/price"
)

const DefaultPriceValue = 10
const DefaultPricePeriodDurationMinutes = 60

func DefaultPrice() *price2.Price {
	return price2.NewPrice(
		price2.NewPricePeriod(0, DefaultPricePeriodDurationMinutes),
		DefaultPriceValue,
		price2.USD(),
	)
}
