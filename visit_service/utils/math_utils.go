package utils

import "math"

func MaxInt64(a, b int64) int64 {
	if a < b {
		return b
	}

	return a
}

func FloorFloat64ToInt(value float64) int {
	floor := math.Floor(value)

	return int(floor)
}
