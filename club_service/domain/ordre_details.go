package domain

type OrderDetails struct {
	sale        int64 //todo change to objects
	certificate int64
	events      []*Event
}

func (o OrderDetails) Events() []*Event {
	return o.events
}

func (o OrderDetails) Sale() int64 {
	return o.sale
}

func (o OrderDetails) Certificate() int64 {
	return o.certificate
}
