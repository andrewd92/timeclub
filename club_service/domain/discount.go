package domain

type Discount struct {
	id    int64
	value float32
}

func (d Discount) Id() int64 {
	return d.id
}

func (d Discount) Value() float32 {
	return d.value
}

func (d Discount) multiplier() float32 {
	return d.value / 100
}

func NewDiscount(id int64, value float32) *Discount {
	return &Discount{id: id, value: value}
}
