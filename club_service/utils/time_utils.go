package utils

import "time"

func MaxTime(a, b time.Time) time.Time {
	if a.Unix() < b.Unix() {
		return b
	}

	return a
}

func MinTime(a, b time.Time) time.Time {
	if a.Unix() > b.Unix() {
		return b
	}

	return a
}
