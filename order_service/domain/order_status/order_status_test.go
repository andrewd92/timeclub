package order_status

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FromString(t *testing.T) {
	cases := []struct {
		Input    string
		Expected OrderStatus
	}{
		{Input: "OPEN", Expected: Open},
		{Input: "PAID", Expected: Paid},
		{Input: "CANCEL", Expected: Cancel},
	}

	for _, testCase := range cases {
		actual, err := FromString(testCase.Input)
		assert.Nil(t, err)
		assert.Equal(t, testCase.Expected, actual)
	}

	orderStatus, err := FromString("TEST")
	assert.Equal(t, None, orderStatus)
	assert.Error(t, err)
}
