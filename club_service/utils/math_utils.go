package utils

func MaxInt64(a, b int64) int64 {
	if a < b {
		return b
	}

	return a
}
