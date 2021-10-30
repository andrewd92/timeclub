package visit

import (
	"github.com/andrewd92/timeclub/visit_service/domain/discount"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	"github.com/andrewd92/timeclub/visit_service/domain/price_list"
	"github.com/andrewd92/timeclub/visit_service/domain/visit/visit_period"
	"math"
	"time"
)

type Visit struct {
	id           int64
	start        *time.Time
	clubId       int64
	orderDetails order_details.OrderDetails
	comment      string
	cardId       int64
	clientName   string
}

func NewVisit(
	id int64,
	start *time.Time,
	clubId int64,
	orderDetails order_details.OrderDetails,
	comment string,
	cardId int64,
	clientName string,
) *Visit {
	return &Visit{
		id:           id,
		start:        start,
		clubId:       clubId,
		orderDetails: orderDetails,
		comment:      comment,
		cardId:       cardId,
		clientName:   clientName,
	}
}

func (v Visit) CalculatePrice(priceList *price_list.PriceList, visitEnd time.Time, openTime string, cardDiscount discount.Discount) (float32, error) {
	visitPeriods, splitErr := v.Period(visitEnd).Split(openTime)

	if nil != splitErr {
		return 0, splitErr
	}

	var result float32 = 0

	for _, period := range visitPeriods {
		result += period.CalculatePrice(priceList, v.orderDetails)
	}

	result -= cardDiscount.From(result)

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

func (v Visit) ClientName() string {
	return v.clientName
}

func (v Visit) ClubId() int64 {
	return v.clubId
}

func (v Visit) CardId() int64 {
	return v.cardId
}
