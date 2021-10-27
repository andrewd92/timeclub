package discount

type Discount struct {
	value float32
}

func (d Discount) Value() float32 {
	return d.value
}

func (d Discount) Multiplier() float32 {
	return d.value / 100
}

func (d Discount) From(price float32) float32 {
	return price * d.Multiplier()
}

func NewDiscount(value float32) *Discount {
	return &Discount{value: value}
}
