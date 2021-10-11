package domain

type Currency struct {
	id        int64
	name      string
	shortName string
}

func NewCurrency(id int64, name string, shortName string) *Currency {
	return &Currency{
		id:        id,
		name:      name,
		shortName: shortName,
	}
}
