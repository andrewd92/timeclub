package event

type eventJson struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Tag      string  `json:"tag"`
	Discount float32 `json:"discount"`
	Start    string  `json:"start"`
	End      string  `json:"end"`
	Price    float32 `json:"price"`
}

func (e Event) Marshal() interface{} {
	return eventJson{
		Id:       e.id,
		Name:     e.name,
		Tag:      e.tag,
		Discount: e.discount.Value(),
		Price:    e.price,
		Start:    e.period.Start().Format("2006-01-02 15:04:05"),
		End:      e.period.End().Format("2006-01-02 15:04:05"),
	}
}

func MarshalAll(events []*Event) []interface{} {
	result := make([]interface{}, len(events))

	for i, event := range events {
		result[i] = event.Marshal()
	}

	return result
}
