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
