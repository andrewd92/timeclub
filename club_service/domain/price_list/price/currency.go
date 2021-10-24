package price

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

func (c Currency) Id() int64 {
	return c.id
}

func (c Currency) Name() string {
	return c.name
}

func (c Currency) ShortName() string {
	return c.shortName
}
