package client_service

import (
	context "context"
	"github.com/andrewd92/timeclub/client_service/api"
	"github.com/andrewd92/timeclub/client_service/domain/client"
	"github.com/andrewd92/timeclub/client_service/infrastructure/repository/client_repository"
)

var server *ClientServerImpl

func Instance() (*ClientServerImpl, error) {
	if nil == server {
		clientRepository, err := client_repository.Instance()

		if err != nil {
			return nil, err
		}

		server = &ClientServerImpl{
			clientRepository: clientRepository,
		}
	}

	return server, nil
}

type ClientServerImpl struct {
	clientRepository client.Repository
}

func (c ClientServerImpl) GetById(_ context.Context, request *api.Request) (*api.Client, error) {
	clientModel, err := c.clientRepository.GetById(request.GetId())

	if err != nil {
		return nil, err
	}

	return &api.Client{
		Id:           clientModel.Id(),
		Name:         clientModel.Name(),
		SecondName:   clientModel.SecondName(),
		Phone:        clientModel.Phone(),
		Email:        clientModel.Email(),
		Birthday:     clientModel.Birthday().Format("2006-01-02"),
		Foto:         clientModel.Foto(),
		Sex:          int32(clientModel.Sex()),
		ClubId:       clientModel.ClubId(),
		City:         clientModel.City(),
		Comment:      clientModel.Comment(),
		Registration: clientModel.Registration().Format("2006-01-02 15:04:05"),
		BonusBalance: clientModel.BonusBalance(),
		CardId:       clientModel.CardId(),
	}, nil
}
