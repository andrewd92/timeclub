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
		Discount: c.discount.Value(),
		Name:     c.name,
		ClubId:   c.ClubId(),
	}
}
