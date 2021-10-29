package card_template

import (
	discountPkg "github.com/andrewd92/timeclub/club_service/domain/discount"
)

type CardTemplate struct {
	id       int64
	discount *discountPkg.Discount
	name     string
	clubId   int64
}

func (ct CardTemplate) Id() int64 {
	return ct.id
}

func (ct CardTemplate) Discount() *discountPkg.Discount {
	return ct.discount
}

func (ct CardTemplate) Name() string {
	return ct.name
}

func (ct CardTemplate) ClubId() int64 {
	return ct.clubId
}

func (ct CardTemplate) WithId(id int64) *CardTemplate {
	return &CardTemplate{id: id, discount: ct.discount, name: ct.name, clubId: ct.clubId}
}

func NewCardTemplate(discount *discountPkg.Discount, name string, clubId int64) *CardTemplate {
	return &CardTemplate{discount: discount, name: name, clubId: clubId}
}
