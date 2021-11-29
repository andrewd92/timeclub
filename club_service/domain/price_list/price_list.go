package price_list

import (
	pricePkg "github.com/andrewd92/timeclub/club_service/domain/price_list/price"
)

type PriceList struct {
	id     int64
	prices []*pricePkg.Price
}

func (p PriceList) Id() int64 {
	return p.id
}

func (p PriceList) Prices() []*pricePkg.Price {
	return p.prices
}

func (p PriceList) WithId(id int64) *PriceList {
	return &PriceList{id: id, prices: p.prices}
}

func NewPriceList(prices []*pricePkg.Price) *PriceList {
	return &PriceList{prices: prices}
}

func NewPriceListWithId(id int64, prices []*pricePkg.Price) *PriceList {
	return &PriceList{id: id, prices: prices}
}
