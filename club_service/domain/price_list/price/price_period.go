package price

// PricePeriod use for configuration price in minutes.
// f.e. for PricePeriod from 0 minute to 60 minute Price will be 10$ per minute
type PricePeriod struct {
	from int
	to   int
}

func NewPricePeriod(from int, to int) *PricePeriod {
	return &PricePeriod{
		from: from,
		to:   to,
	}
}

func (p PricePeriod) From() int {
	return p.from
}

func (p PricePeriod) To() int {
	return p.to
}
