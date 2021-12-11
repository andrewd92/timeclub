package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPay(t *testing.T) {
	order := DefaultOrder()

	order.Pay()

	assert.Equal(t, Paid, order.Status)

	order.Cancel()
	assert.Equal(t, Paid, order.Status)
}

func TestCancel(t *testing.T) {
	order := DefaultOrder()

	order.Cancel()

	assert.Equal(t, Cancel, order.Status)
	order.Pay()
	assert.Equal(t, Cancel, order.Status)
}

func DefaultOrder() *Order {
	return &Order{
		Id:     1,
		Status: Open,
		Start:  time.Now(),
		End:    nil,
		Visits: []int64{1, 5},
	}
}
