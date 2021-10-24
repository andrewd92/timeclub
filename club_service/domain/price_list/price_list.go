package price_list

import (
	pricePkg "github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	"sort"
)

type PriceList struct {
	prices []*pricePkg.Price
	max    float32
}

func (p PriceList) Calculate(durationMinutes int) float32 {
	var total float32 = 0

	for _, price := range p.prices {
		total += price.Calculate(durationMinutes)
	}

	return total
}

func (p PriceList) Prices() []*pricePkg.Price {
	prices := make([]*pricePkg.Price, len(p.prices))
	copy(prices, p.prices)
	return prices
}

func (p PriceList) Max() float32 {
	return p.max
}

func NewPriceList(prices []*pricePkg.Price) *PriceList {
	sort.Slice(prices, func(i, j int) bool {
		if prices[i].PricePeriod().From() > prices[j].PricePeriod().From() {
			return false
		}

		return true
	})

	var max float32 = 0

	for _, price := range prices {
		max += price.Max()
	}

	return &PriceList{prices: prices, max: max}
}
