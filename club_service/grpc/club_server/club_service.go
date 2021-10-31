package club_server

import (
	context "context"
	"github.com/andrewd92/timeclub/club_service/api"
	clubPkg "github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/club_repository"
)

var server *ClubServerImpl

func Instance() *ClubServerImpl {
	if nil == server {
		server = &ClubServerImpl{
			clubRepository: club_repository.Instance(),
		}
	}

	return server
}

type ClubServerImpl struct {
	clubRepository clubPkg.Repository
}

func (c ClubServerImpl) GetById(ctx context.Context, request *api.Request) (*api.Club, error) {
	club, err := c.clubRepository.GetById(request.Id)

	if err != nil {
		return nil, err
	}

	priceList := club.PriceList()
	prices := make([]*api.Price, 0, len(priceList.Prices()))

	for _, price := range priceList.Prices() {
		pricePeriod := map[string]int32{
			"from": int32(price.PricePeriod().From()),
			"to":   int32(price.PricePeriod().To()),
		}
		prices = append(prices, &api.Price{
			PricePeriod:    pricePeriod,
			ValuePerMinute: price.ValuePerMinute(),
			Currency:       price.Currency().ShortName(),
		})
	}

	return &api.Club{
		Id:       club.Id(),
		Name:     club.Name(),
		OpenTime: club.OpenTime(),
		Currency: club.Currency().ShortName(),
		Prices:   prices,
	}, nil
}
