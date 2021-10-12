package domain

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

func NewTimePeriod(start time.Time, end time.Time) *TimePeriod {
	return &TimePeriod{start: start, end: end}
}
