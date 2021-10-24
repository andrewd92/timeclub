package time_period

import (
	"github.com/andrewd92/timeclub/club_service/utils"
	"time"
)

type TimePeriod struct {
	start time.Time
	end   time.Time
}

func (t TimePeriod) CommonSeconds(otherPeriod *TimePeriod) int64 {
	start := utils.MaxTime(t.start, otherPeriod.Start()).Unix()
	end := utils.MinTime(t.end, otherPeriod.End()).Unix()

	return utils.MaxInt64(0, end-start)
}

func (t TimePeriod) Start() time.Time {
	return t.start
}

func (t TimePeriod) End() time.Time {
	return t.end
}

func (t TimePeriod) Duration() int {
	return int(t.end.Unix() - t.start.Unix())
}

func (t TimePeriod) DurationMinutes() int {
	duration := float64(t.Duration()) / 60.0
	return utils.FloorFloat64ToInt(duration)
}

func NewTimePeriod(start time.Time, end time.Time) *TimePeriod {
	return &TimePeriod{start: start, end: end}
}
