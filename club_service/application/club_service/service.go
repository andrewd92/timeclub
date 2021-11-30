package club_service

import (
	"errors"
	"github.com/andrewd92/timeclub/club_service/api/http/create"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	currencyPkg "github.com/andrewd92/timeclub/club_service/domain/currency"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	log "github.com/sirupsen/logrus"
)

type ClubService interface {
	Create(request create.Request) (interface{}, error)
}

type clubServiceImpl struct {
	clubRepository      club.Repository
	priceListRepository price_list.Repository
	currencyRepository  currencyPkg.Repository
}

func (s clubServiceImpl) Create(request create.Request) (interface{}, error) {
	priceList := createPriceList(request)
	priceList, err := s.priceListRepository.Save(priceList)
	if err != nil {
		log.WithError(err).WithField("price_list", priceList).Error("can not store price list")
		return nil, errors.New("db error")
	}

	currency, err := s.currencyRepository.GetById(request.CurrencyId)
	if err != nil {
		log.WithError(err).WithField("currency", request.CurrencyId).Error("can not find currency")
		return nil, errors.New("db error")
	}

	clubEntity := club.NewClub(request.Name, request.OpenTime, priceList, currency)

	newClub, err := s.clubRepository.Save(clubEntity)
	if err != nil {
		log.WithError(err).WithField("club", clubEntity).Error("can not store clubEntity")
		return nil, errors.New("db error")
	}

	return newClub.Marshal(), nil
}

func createPriceList(request create.Request) *price_list.PriceList {
	prices := make([]*price.Price, len(request.PriceList))

	for i, priceModel := range request.PriceList {
		newPrice := price.NewPrice(
			price.NewPricePeriod(priceModel.PricePeriod.From, priceModel.PricePeriod.To),
			priceModel.ValuePerMinute,
		)
		prices[i] = newPrice
	}

	return price_list.NewPriceList(prices)
}
