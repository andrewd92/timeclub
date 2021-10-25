package club

type clubJson struct {
	Id       int64         `json:"id"`
	Name     string        `json:"name"`
	OpenTime string        `json:"open_time"`
	Currency string        `json:"currency"`
	Prices   []interface{} `json:"prices"`
}

func (c Club) Marshal() interface{} {
	return clubJson{
		Id:       c.id,
		Name:     c.name,
		OpenTime: c.openTime,
		Currency: c.currency.Name(),
		Prices:   c.PriceList().Marshall(),
	}
}

func MarshalAll(clubs []*Club) []interface{} {
	result := make([]interface{}, len(clubs))

	for i, club := range clubs {
		result[i] = club.Marshal()
	}

	return result
}
