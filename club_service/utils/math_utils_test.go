package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxInt64ShouldReturnMaxValue(t *testing.T) {
	var expected int64 = 100
	assert.Equal(t, expected, MaxInt64(expected, 99))
	assert.Equal(t, expected, MaxInt64(99, expected))
}

func TestFloorFloat64ToInt(t *testing.T) {
	floatValue := 3.99

	assert.Equal(t, 3, FloorFloat64ToInt(floatValue))
}
