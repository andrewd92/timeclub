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

func (v *VisitPeriod) split(closeTime string) ([]*VisitPeriod, error) {
	closeHour, closeMinute, splitErr := utils.SplitTimeString(closeTime)
	if splitErr != nil {
		return nil, splitErr
	}

	oneDayDuration, _ := time.ParseDuration("24h")
	oneDayDurationSec := int64(oneDayDuration.Seconds())

	closeTs := time.Date(v.start.Year(), v.start.Month(), v.start.Day(), closeHour, closeMinute, 0, 0, v.start.Location()).Unix()
	startTs := v.start.Unix()
	endTs := v.end.Unix()

	if startTs >= closeTs {
		closeTs += oneDayDurationSec
	}

	if endTs <= closeTs {
		period := *v
		return []*VisitPeriod{&period}, nil
	}

	periodsTs := []int64{startTs}

	for endTs > closeTs {
		periodsTs = append(periodsTs, closeTs)
		closeTs += oneDayDurationSec
	}

	periodsTs = append(periodsTs, endTs)

	periods := make([]*VisitPeriod, len(periodsTs)-1)

	for i := 0; i < len(periodsTs)-1; i++ {
		periodStartTs := periodsTs[i]
		startTime := time.Unix(periodStartTs, 0)
		endTime := time.Unix(periodsTs[i+1], 0)
		firstMinute := int((periodStartTs - startTs) / 60)

		periods = append(periods, NewVisitPeriodFromMinute(startTime, endTime, firstMinute))
	}

	return periods, nil
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

func NewVisitPeriodFromMinute(start time.Time, end time.Time, firstMinute int) *VisitPeriod {
	return &VisitPeriod{start: start, end: end, firstMinute: firstMinute}
}

func NewVisitPeriod(start time.Time, end time.Time) *VisitPeriod {
	return &VisitPeriod{start: start, end: end, firstMinute: 0}
}
