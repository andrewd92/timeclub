package price_list

import (
	"github.com/andrewd92/timeclub/visit_service/domain/price_list/price"
)

func DefaultPriceList() *PriceList {
	return NewPriceList([]*price.Price{price.DefaultPrice()})
}
