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

type Marshaller interface {
	MarshalAll(visits []*Visit, club *api.Club, now time.Time) ([]interface{}, error)
	Marshal(visit *Visit, now time.Time, club *api.Club) (interface{}, error)
}

type MarshallerImpl struct {
}

func (m MarshallerImpl) Marshal(visit *Visit, now time.Time, club *api.Club) (interface{}, error) {
	price, err := visit.CalculatePrice(club, now, *discount.NewDiscount(0))

	if err != nil {
		return nil, err
	}

	return visitJson{
		Id:         visit.id,
		Start:      visit.start.Format("2006-01-02 15:04:05"),
		ClientName: visit.clientName,
		ClubId:     visit.ClubId(),
		Comment:    visit.comment,
		Price:      price,
		Currency:   club.Currency.ShortName,
		Duration:   visit.Duration(now),
		CardId:     visit.CardId(),
	}, nil
}

func (m MarshallerImpl) MarshalAll(visits []*Visit, club *api.Club, now time.Time) ([]interface{}, error) {
	result := make([]interface{}, len(visits))

	for i, visit := range visits {

		json, err := m.Marshal(visit, now, club)

		if err != nil {
			return nil, err
		}

		result[i] = json
	}

	return result, nil
}
