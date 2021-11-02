package card

type cardJson struct {
	Id       int64   `json:"id"`
	Discount float32 `json:"discount"`
	Name     string  `json:"name"`
	ClubId   int64   `json:"club_id"`
}

func (c Card) Marshal() interface{} {
	return cardJson{
		Id:       c.id,
		Discount: float32(c.discount),
		Name:     c.name,
		ClubId:   c.ClubId(),
	}
}

func MarshalAll(cards []*Card) []interface{} {
	result := make([]interface{}, 0, len(cards))

	for _, card := range cards {
		result = append(result, card.Marshal())
	}

	return result
}
