package card

import "github.com/andrewd92/timeclub/card_service/domain/discount"

type Card struct {
	id       int64
	discount discount.Discount
	name     string
	clubId   int64
}

func NewCard(discount discount.Discount, name string, clubId int64) *Card {
	return &Card{discount: discount, name: name, clubId: clubId}
}

func (c Card) ClubId() int64 {
	return c.clubId
}

func (c Card) Id() int64 {
	return c.id
}

func (c Card) Discount() discount.Discount {
	return c.discount
}

func (c Card) Name() string {
	return c.name
}

func (c Card) WithId(id int64) *Card {
	return &Card{id: id, discount: c.discount, name: c.name, clubId: c.ClubId()}
}
