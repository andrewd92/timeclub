package grpc

import (
	"context"
	"github.com/andrewd92/timeclub/card_service/api"
	cardPkg "github.com/andrewd92/timeclub/card_service/domain/card"
	"github.com/andrewd92/timeclub/card_service/infrastructure/repository/card_repository"
)

var server api.CardServiceServer

func Instance() (api.CardServiceServer, error) {
	if nil == server {
		server = &CardServerImpl{
			repository: card_repository.Instance(),
		}
	}

	return server, nil
}

type CardServerImpl struct {
	repository cardPkg.Repository
}

func (s CardServerImpl) GetById(_ context.Context, request *api.Request) (*api.Card, error) {
	card, err := s.repository.GetById(request.Id)
	if err != nil {
		return nil, err
	}

	return &api.Card{
		Id:       card.Id(),
		Discount: float32(card.Discount()),
		Name:     card.Name(),
		ClubId:   card.ClubId(),
	}, nil
}
