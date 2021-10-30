package price

const DefaultPriceValue = 10
const DefaultPricePeriodDurationMinutes = 60

func DefaultPrice() *Price {
	return NewPrice(
		NewPricePeriod(0, DefaultPricePeriodDurationMinutes),
		DefaultPriceValue,
		USD(),
	)
}
