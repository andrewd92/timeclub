package domain

type Club struct {
	id        int64
	name      string
	openTime  string
	priceList *PriceList
	currency  *Currency
}

func NewClub(id int64, name string, openTime string, priceList *PriceList, currency *Currency) *Club {
	return &Club{id: id, name: name, openTime: openTime, priceList: priceList, currency: currency}
}

func (c Club) Currency() *Currency {
	return c.currency
}

func (c Club) PriceList() *PriceList {
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
