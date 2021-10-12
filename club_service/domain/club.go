package domain

type Club struct {
	id          int64
	name        string
	openTime    string
	priceMatrix []*Price
}

func NewClub(id int64, name string, openTime string, priceMatrix []*Price) *Club {
	return &Club{id: id, name: name, openTime: openTime, priceMatrix: priceMatrix}
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

func (c Club) PriceMatrix() []*Price {
	return c.priceMatrix
}
