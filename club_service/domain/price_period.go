package domain

// PricePeriod use for configuration price in minutes.
// f.e. for PricePeriod from 0 minute to 60 minute Price will be 10$ per minute
type PricePeriod struct {
	from int64
	to   int64
}

func NewPricePeriod(from int64, to int64) *PricePeriod {
	return &PricePeriod{
		from: from,
		to:   to,
	}
}

func (p PricePeriod) From() int64 {
	return p.from
}

func (p PricePeriod) To() int64 {
	return p.to
}

func (p *PricePeriod) TimeForPay(durationMinutes int64) int64 {
	if p.from > durationMinutes {
		return 0
	}

	if p.to > durationMinutes {
		return durationMinutes - p.from
	}

	return p.totalTime()
}

func (p PricePeriod) totalTime() int64 {
	return p.to - p.from
}
