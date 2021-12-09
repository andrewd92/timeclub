package domain

import (
	"github.com/andrewd92/timeclub/order_service/domain/order_status"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPay(t *testing.T) {
	order := DefaultOrder()

	order.Pay()

	assert.Equal(t, order_status.Paid, order.Status)

	order.Cancel()
	assert.Equal(t, order_status.Paid, order.Status)
}

func TestCancel(t *testing.T) {
	order := DefaultOrder()

	order.Cancel()

	assert.Equal(t, order_status.Cancel, order.Status)
	order.Pay()
	assert.Equal(t, order_status.Cancel, order.Status)
}

func DefaultOrder() *Order {
	return &Order{
		Id:     1,
		Status: order_status.Open,
		Start:  time.Now(),
		End:    nil,
		Visits: []int64{1, 5},
	}
}
