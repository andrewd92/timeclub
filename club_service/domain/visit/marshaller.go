package visit

import "time"

type visitJson struct {
	Id         int64   `json:"id"`
	Start      string  `json:"start"`
	ClientId   int64   `json:"client_id"`
	ClientName string  `json:"client_name"`
	ClubId     int64   `json:"club_id"`
	Comment    string  `json:"comment"`
	Price      float32 `json:"price"`
	Currency   string  `json:"currency"`
	Duration   int     `json:"duration"`
}

func (v Visit) Marshal(now time.Time) (interface{}, error) {
	price, err := v.CalculatePrice(v.club.PriceList(), now)

	if err != nil {
		return nil, err
	}

	return visitJson{
		Id:         v.id,
		Start:      v.start.Format("2006-01-02 15:04:05"),
		ClientId:   v.client.Id(),
		ClientName: v.client.Name(),
		ClubId:     v.club.Id(),
		Comment:    v.comment,
		Price:      price,
		Currency:   v.club.Currency().ShortName(),
		Duration:   v.Duration(now),
	}, nil
}

func MarshalAll(visits []*Visit) ([]interface{}, error) {
	result := make([]interface{}, len(visits))

	now := time.Now()

	for i, visit := range visits {

		json, err := visit.Marshal(now)

		if err != nil {
			return nil, err
		}

		result[i] = json
	}

	return result, nil
}
