package domain

import "sort"

type PriceList struct {
	prices []*Price
	max    float32
}

func (p PriceList) Calculate(visitPeriod *VisitPeriod) float32 {
	var total float32 = 0

	for _, price := range p.prices {
		total += price.Calculate(visitPeriod)
	}

	return total
}

func (p PriceList) Prices() []*Price {
	prices := make([]*Price, len(p.prices))
	copy(prices, p.prices)
	return prices
}

func (p PriceList) Max() float32 {
	return p.max
}

func NewPriceList(prices []*Price) *PriceList {
	sort.Slice(prices, func(i, j int) bool {
		if prices[i].period.From() > prices[j].period.From() {
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
