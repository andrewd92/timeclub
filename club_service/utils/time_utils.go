package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

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

/**
Split time to hour and minute.
F.e. splitting string "12:34" will return (12, 34, nil)
*/
func SplitTimeString(timeString string) (int, int, error) {
	splitTime := strings.Split(timeString, ":")

	if len(splitTime) != 2 {
		return 0, 0, errors.New("Split openTime should have format H:m, given: " + timeString)
	}

	hour, parseHourError := strconv.ParseInt(splitTime[0], 10, 64)
	if parseHourError != nil {
		return 0, 0, errors.New("Can not parse hour to int. Hour: " + splitTime[0])
	}

	minute, parseMinuteError := strconv.ParseInt(splitTime[1], 10, 64)
	if parseMinuteError != nil {
		return 0, 0, errors.New("Can not parse minute to int. Minute: " + splitTime[0])
	}

	return int(hour), int(minute), nil
}
