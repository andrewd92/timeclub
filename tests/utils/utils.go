package utils

const TimeFormat = "2006-01-02 15:04:05"

func ParseInt64(data interface{}) int64 {
	return int64(data.(float64))
}
