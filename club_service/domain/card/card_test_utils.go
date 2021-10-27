package card

import "github.com/andrewd92/timeclub/club_service/domain/discount"

func DefaultCard() *Card {
	return NewCard(1, discount.NewDiscount(10.0), "Best Client", 1)
}
