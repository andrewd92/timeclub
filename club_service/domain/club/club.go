package club

import (
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
)

type Club struct {
	id        int64
	name      string
	openTime  string
	priceList *price_list.PriceList
	currency  *price.Currency
}

func NewClub(id int64, name string, openTime string, priceList *price_list.PriceList, currency *price.Currency) *Club {
	return &Club{id: id, name: name, openTime: openTime, priceList: priceList, currency: currency}
}

func (c Club) Currency() *price.Currency {
	return c.currency
}

func (c Club) PriceList() *price_list.PriceList {
	return c.priceList
}

func (c Club) OpenTime() string {
	return c.openTime
}

func (c Club) Id() int64 {
	return c.id
}

func (c Club) Name() string {
	return c.name
}
