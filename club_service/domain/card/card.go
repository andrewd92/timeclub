package card

import discountPkg "github.com/andrewd92/timeclub/club_service/domain/discount"

type Card struct {
	id       int64
	discount discountPkg.Discount
	name     string
	cafeId   int64
}

func NewCard(id int64, discount discountPkg.Discount, name string, cafeId int64) *Card {
	return &Card{id: id, discount: discount, name: name, cafeId: cafeId}
}

func (c Card) CafeId() int64 {
	return c.cafeId
}

func (c Card) Id() int64 {
	return c.id
}

func (c Card) Discount() discountPkg.Discount {
	return c.discount
}

func (c Card) Name() string {
	return c.name
}
