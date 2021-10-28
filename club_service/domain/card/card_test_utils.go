package card

import "github.com/andrewd92/timeclub/club_service/domain/discount"

func DefaultCard() *Card {
	return DefaultCardWithId(1)
}

func DefaultCardWithId(id int64) *Card {
	return NewCard(id, discount.NewDiscount(10.0), "Best Client", 1)
}
