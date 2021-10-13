package domain

type Event struct {
	name     string
	tag      string
	discount Discount
	period   TimePeriod
	price    float32
}

func (e Event) Name() string {
	return e.name
}

func (e Event) Tag() string {
	return e.tag
}

func (e Event) Discount() Discount {
	return e.discount
}

func (e Event) Period() TimePeriod {
	return e.period
}

func NewEvent(name string, tag string, discount Discount, period TimePeriod) *Event {
	return &Event{name: name, tag: tag, discount: discount, period: period}
}
