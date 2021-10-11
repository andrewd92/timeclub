package domain

type Club struct {
	id          int64
	name        string
	priceMatrix []*Price
}

func NewClub(id int64, name string, priceMatrix []*Price) *Club {
	return &Club{
		id:          id,
		name:        name,
		priceMatrix: priceMatrix,
	}
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
