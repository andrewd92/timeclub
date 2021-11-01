package visit_service

import (
	"github.com/andrewd92/timeclub/club_service/api"
	"github.com/andrewd92/timeclub/visit_service/client/club_service"
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	"github.com/andrewd92/timeclub/visit_service/domain/price_list"
	pricePkg "github.com/andrewd92/timeclub/visit_service/domain/price_list/price"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"time"
)

type VisitService interface {
	Create(clubId int64, cardId int64) (interface{}, error)
}

type visitServiceImpl struct {
	visitRepository visit.Repository
}

func (v visitServiceImpl) Create(clubId int64, cardId int64) (interface{}, error) {
	now := time.Now()

	newVisit := visit.NewVisit(
		&now,
		clubId,
		order_details.NewOrderDetails(make([]*event.Event, 0)),
		"",
		cardId,
		"Guest",
	)

	visitModel, saveErr := v.visitRepository.Save(newVisit)

	if nil != saveErr {
		return nil, saveErr
	}

	club, clubServiceErr := club_service.GetById(clubId)
	if nil != clubServiceErr {
		return nil, clubServiceErr
	}

	priceList := createPriceList(club)

	marshal, marshalErr := visitModel.Marshal(now.Add(time.Hour), priceList, pricePkg.USD())
	if marshalErr != nil {
		return nil, marshalErr
	}

	return marshal, nil
}

func createPriceList(club *api.Club) *price_list.PriceList {
	prices := make([]*pricePkg.Price, 0, len(club.Prices))

	for _, priceResponse := range club.Prices {
		pricePeriod := pricePkg.NewPricePeriod(int(priceResponse.PricePeriod.From), int(priceResponse.PricePeriod.To))
		price := pricePkg.NewPrice(pricePeriod, priceResponse.ValuePerMinute, pricePkg.USD())
		prices = append(prices, price)
	}

	return price_list.NewPriceList(prices)
}
