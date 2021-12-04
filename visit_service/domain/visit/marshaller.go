package visit

import (
	"github.com/andrewd92/timeclub/club_service/api"
	"github.com/andrewd92/timeclub/visit_service/domain/discount"
	"time"
)

type visitJson struct {
	Id         int64   `json:"id"`
	Start      string  `json:"start"`
	ClientName string  `json:"client_name"`
	ClubId     int64   `json:"club_id"`
	Comment    string  `json:"comment"`
	Price      float32 `json:"price"`
	Currency   string  `json:"currency"`
	Duration   int     `json:"duration"`
	CardId     int64   `json:"card_id"`
}

func (v Visit) Marshal(now time.Time, club *api.Club) (interface{}, error) {
	price, err := v.CalculatePrice(club, now, *discount.NewDiscount(10.0))

	if err != nil {
		return nil, err
	}

	return visitJson{
		Id:         v.id,
		Start:      v.start.Format("2006-01-02 15:04:05"),
		ClientName: v.clientName,
		ClubId:     v.ClubId(),
		Comment:    v.comment,
		Price:      price,
		Currency:   club.Currency.ShortName,
		Duration:   v.Duration(now),
		CardId:     v.CardId(),
	}, nil
}

func MarshalAll(visits []*Visit, club *api.Club) ([]interface{}, error) {
	result := make([]interface{}, len(visits))

	now := time.Now()

	for i, visit := range visits {

		json, err := visit.Marshal(now, club)

		if err != nil {
			return nil, err
		}

		result[i] = json
	}

	return result, nil
}
