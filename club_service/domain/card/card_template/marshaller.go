package card_template

type cardTemplateJson struct {
	Id       int64   `json:"id"`
	Discount float32 `json:"discount"`
	Name     string  `json:"name"`
	ClubId   int64   `json:"club_id"`
}

func (ct CardTemplate) Marshal() interface{} {
	return cardTemplateJson{
		Id:       ct.id,
		Discount: ct.discount.Value(),
		Name:     ct.name,
		ClubId:   ct.ClubId(),
	}
}

func MarshalAll(templates []*CardTemplate) []interface{} {
	result := make([]interface{}, 0, len(templates))

	for _, template := range templates {
		result = append(result, template.Marshal())
	}

	return result
}
