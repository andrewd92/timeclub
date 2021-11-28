package price_list

import (
	pricePkg "github.com/andrewd92/timeclub/club_service/domain/price_list/price"
)

type PriceList struct {
	prices []*pricePkg.Price
}

func (p PriceList) Prices() []*pricePkg.Price {
	return p.prices
}

func NewPriceList(prices []*pricePkg.Price) *PriceList {
	return &PriceList{prices: prices}
}

func Empty() *PriceList {
	prices := make([]*pricePkg.Price, 0)

	return &PriceList{prices: prices}
}
