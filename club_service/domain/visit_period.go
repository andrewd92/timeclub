package domain

import (
	"github.com/andrewd92/timeclub/club_service/utils"
	"time"
)

type VisitPeriod struct {
	start       time.Time
	end         time.Time
	firstMinute int
}

func (v *VisitPeriod) Split(closeTime string) ([]*VisitPeriod, error) {
	closeTs, closeTsCalculationErr := v.calculateCloseTs(closeTime)
	if closeTsCalculationErr != nil {
		return nil, closeTsCalculationErr
	}

	if v.end.Unix() <= closeTs {
		period := *v
		return []*VisitPeriod{&period}, nil
	}

	tsRange := v.splitToTsRange(closeTs)

	return v.calculatePeriodFromTsRange(tsRange), nil
}

func (v VisitPeriod) Start() time.Time {
	return v.start
}

func (v VisitPeriod) End() time.Time {
	return v.end
}

func (v VisitPeriod) FirstMinute() int {
	return v.firstMinute
}

func (v VisitPeriod) calculateCloseTs(closeTime string) (int64, error) {
	closeHour, closeMinute, splitErr := utils.SplitTimeString(closeTime)
	if splitErr != nil {
		return 0, splitErr
	}

	closeTs := time.Date(v.start.Year(), v.start.Month(), v.start.Day(), closeHour, closeMinute, 0, 0, v.start.Location()).Unix()

	if v.start.Unix() >= closeTs {
		closeTs += utils.OneDayDuration
	}

	return closeTs, nil
}

func (v VisitPeriod) splitToTsRange(closeTs int64) []int64 {
	endTs := v.end.Unix()
	tsRange := []int64{v.start.Unix()}

	for endTs > closeTs {
		tsRange = append(tsRange, closeTs)
		closeTs += utils.OneDayDuration
	}

	tsRange = append(tsRange, endTs)

	return tsRange
}

func (v VisitPeriod) calculatePeriodFromTsRange(tsRange []int64) []*VisitPeriod {
	var periods []*VisitPeriod

	for i := 0; i < len(tsRange)-1; i++ {
		periodStartTs := tsRange[i]
		startTime := time.Unix(periodStartTs, 0)
		endTime := time.Unix(tsRange[i+1], 0)
		firstMinute := int((periodStartTs - v.start.Unix()) / 60)

		periods = append(periods, NewVisitPeriodFromMinute(startTime, endTime, firstMinute))
	}

	return periods
}

func NewVisitPeriodFromMinute(start time.Time, end time.Time, firstMinute int) *VisitPeriod {
	return &VisitPeriod{start: start, end: end, firstMinute: firstMinute}
}

func NewVisitPeriod(start time.Time, end time.Time) *VisitPeriod {
	return &VisitPeriod{start: start, end: end, firstMinute: 0}
}
