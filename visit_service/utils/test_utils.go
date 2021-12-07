package utils

import (
	"github.com/andrewd92/timeclub/club_service/api"
	"time"
)

func DefaultClub() *api.Club {
	openTime := time.Now().Add(-12 * time.Hour).Format("15:04")

	return &api.Club{
		Id: 1, Name: "Test Club", OpenTime: openTime, Prices: []*api.Price{DefaultPrice()}, Currency: USD(),
	}
}

func ClubWithWrongOpenTime() *api.Club {
	return &api.Club{
		Id: 1, Name: "Test Club", OpenTime: "12:00:00", Prices: []*api.Price{DefaultPrice()},
	}
}

const DefaultPriceValue = 10
const DefaultPricePeriodDurationMinutes = 360

func DefaultPrice() *api.Price {
	return &api.Price{
		PricePeriod: &api.PricePeriod{
			From: 0,
			To:   DefaultPricePeriodDurationMinutes,
		},
		ValuePerMinute: DefaultPriceValue,
	}
}

func USD() *api.Currency {
	return &api.Currency{
		Name:      "USD",
		ShortName: "$",
	}
}
