package card

type cardJson struct {
	Id       int64       `json:"id"`
	Discount float32     `json:"discount"`
	Name     string      `json:"name"`
	ClubId   interface{} `json:"club_id"`
}

func (c Card) Marshal() interface{} {
	var clubId interface{} = nil
	if c.ClubId() != 0 {
		clubId = c.ClubId()
	}
	return cardJson{
		Id:       c.id,
		Discount: float32(c.discount),
		Name:     c.name,
		ClubId:   clubId,
	}
}

func MarshalAll(cards []*Card) []interface{} {
	result := make([]interface{}, 0, len(cards))

	for _, card := range cards {
		result = append(result, card.Marshal())
	}

	return result
}
