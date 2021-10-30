package price

func USD() *Currency {
	return NewCurrency(1, "USD", "$")
}
