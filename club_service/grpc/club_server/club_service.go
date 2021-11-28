package club_server

import (
	"context"
	"github.com/andrewd92/timeclub/club_service/api"
	clubPkg "github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/club_repository"
)

var server api.ClubServiceServer

func Instance() (api.ClubServiceServer, error) {
	if nil == server {
		clubRepository, err := club_repository.Instance()

		if err != nil {
			return nil, err
		}

		server = &ClubServerImpl{
			clubRepository: clubRepository,
		}
	}

	return server, nil
}

type ClubServerImpl struct {
	clubRepository clubPkg.Repository
}

func (c ClubServerImpl) GetById(_ context.Context, request *api.Request) (*api.Club, error) {
	club, err := c.clubRepository.GetById(request.Id)

	if err != nil {
		return nil, err
	}

	priceList := club.PriceList()
	prices := make([]*api.Price, 0, len(priceList.Prices()))

	for _, price := range priceList.Prices() {
		pricePeriod := &api.PricePeriod{
			From: int32(price.PricePeriod().From()),
			To:   int32(price.PricePeriod().To()),
		}

		prices = append(prices, &api.Price{
			PricePeriod:    pricePeriod,
			ValuePerMinute: price.ValuePerMinute(),
		})
	}

	return &api.Club{
		Id:       club.Id(),
		Name:     club.Name(),
		OpenTime: club.OpenTime(),
		Currency: &api.Currency{Name: club.Currency().Name(), ShortName: club.Currency().ShortName()},
		Prices:   prices,
	}, nil
}
