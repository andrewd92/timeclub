package order_service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create(t *testing.T) {
	testingService := testInstance()
	visits := []int64{1, 2, 4}

	response, err := testingService.Create(visits)

	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"response": "ok"}, response)
}

func testInstance() *OrderServiceImpl {
	return &OrderServiceImpl{}
}
