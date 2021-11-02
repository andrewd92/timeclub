package utils

import "github.com/andrewd92/timeclub/club_service/api"

func DefaultClub() *api.Club {
	return &api.Club{
		Id: 1, Name: "Test Club", OpenTime: "12:00", Prices: []*api.Price{DefaultPrice()},
	}
}

func ClubWithWrongOpenTime() *api.Club {
	return &api.Club{
		Id: 1, Name: "Test Club", OpenTime: "12:00:00", Prices: []*api.Price{DefaultPrice()},
	}
}

const DefaultPriceValue = 10
const DefaultPricePeriodDurationMinutes = 60

func DefaultPrice() *api.Price {
	return &api.Price{
		PricePeriod: &api.PricePeriod{
			From: 0,
			To:   DefaultPricePeriodDurationMinutes,
		},
		ValuePerMinute: DefaultPriceValue,
	}
}
