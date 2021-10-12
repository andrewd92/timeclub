package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMaxTimeShouldReturnMaxTime(t *testing.T) {
	now := time.Now()
	oneSecondDuration, _ := time.ParseDuration("1s")
	after := now.Add(oneSecondDuration)

	assert.Equal(t, after, MaxTime(now, after))
	assert.Equal(t, after, MaxTime(after, now))
}

func TestMinTimeShouldReturnMixTime(t *testing.T) {
	now := time.Now()
	oneSecondDuration, _ := time.ParseDuration("1s")
	after := now.Add(oneSecondDuration)

	assert.Equal(t, now, MinTime(now, after))
	assert.Equal(t, now, MinTime(after, now))
}
