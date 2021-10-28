package visit

import (
	"github.com/andrewd92/timeclub/club_service/domain/card"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/order_details"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/visit/visit_period"
	"math"
	"time"
)

type Visit struct {
	id           int64
	start        *time.Time
	club         *club.Club
	orderDetails order_details.OrderDetails
	comment      string
	card         *card.Card
	clientName   string
}

func NewVisit(
	id int64,
	start *time.Time,
	club *club.Club,
	orderDetails order_details.OrderDetails,
	comment string,
	card *card.Card,
	clientName string,
) *Visit {
	return &Visit{
		id:           id,
		start:        start,
		club:         club,
		orderDetails: orderDetails,
		comment:      comment,
		card:         card,
		clientName:   clientName,
	}
}

func (v Visit) CalculatePrice(priceList *price_list.PriceList, visitEnd time.Time) (float32, error) {
	visitPeriods, splitErr := v.Period(visitEnd).Split(v.club.OpenTime())

	if nil != splitErr {
		return 0, splitErr
	}

	var result float32 = 0

	for _, period := range visitPeriods {
		result += period.CalculatePrice(priceList, v.orderDetails)
	}

	result -= v.card.Discount().From(result)

	return result, nil
}

func (v Visit) Period(endDate time.Time) *visit_period.VisitPeriod {
	return visit_period.NewVisitPeriod(*v.Start(), endDate)
}

func (v *Visit) Duration(endDate time.Time) int {
	durationOfMinutes := endDate.Sub(*v.start).Minutes()

	return int(math.Floor(durationOfMinutes))
}

func (v Visit) OrderDetails() order_details.OrderDetails {
	return v.orderDetails
}

func (v Visit) Comment() string {
	return v.comment
}

func (v Visit) Id() int64 {
	return v.id
}

func (v Visit) Start() *time.Time {
	return v.start
}

func (v Visit) Club() *club.Club {
	return v.club
}

func (v Visit) Card() *card.Card {
	return v.card
}

func (v Visit) ClientName() string {
	return v.clientName
}
