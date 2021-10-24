package price_list

import (
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriceList_PricesShouldReturnCopy(t *testing.T) {
	list := DefaultPriceList()

	prices := list.Prices()
	prices = append(prices, price.DefaultPrice())

	assert.NotEqual(t, prices, list.prices)
}
