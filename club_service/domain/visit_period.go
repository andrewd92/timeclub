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

func (v *VisitPeriod) Split(splitTime string) ([]*VisitPeriod, error) {
	firstSplitTs, firstSplitTsCalculationErr := v.calculateFirstSplitTs(splitTime)
	if firstSplitTsCalculationErr != nil {
		return nil, firstSplitTsCalculationErr
	}

	if v.end.Unix() <= firstSplitTs {
		period := *v
		return []*VisitPeriod{&period}, nil
	}

	tsRange := v.splitToTsRange(firstSplitTs)

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

func (v VisitPeriod) calculateFirstSplitTs(splitTime string) (int64, error) {
	splitHour, splitMinute, splitErr := utils.SplitTimeString(splitTime)
	if splitErr != nil {
		return 0, splitErr
	}

	splitTs := time.Date(v.start.Year(), v.start.Month(), v.start.Day(), splitHour, splitMinute, 0, 0, v.start.Location()).Unix()

	if v.start.Unix() >= splitTs {
		splitTs += utils.OneDayDuration
	}

	return splitTs, nil
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
