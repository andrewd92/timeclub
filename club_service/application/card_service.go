package application

import (
	"github.com/andrewd92/timeclub/club_service/domain/card"
	discountPkg "github.com/andrewd92/timeclub/club_service/domain/discount"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/card_repository"
)

type cardServiceImpl struct {
	cardRepository card.Repository
}

type CardService interface {
	Create(cafeId int64, discount float32, name string) (interface{}, error)
}

var service CardService

func CardServiceInstance() CardService {
	if nil == service {
		service = &cardServiceImpl{
			cardRepository: card_repository.Instance(),
		}
	}

	return service
}

func (s cardServiceImpl) Create(clubId int64, discount float32, name string) (interface{}, error) {
	newCard := card.NewCard(discountPkg.NewDiscount(discount), name, clubId)
	cardModel, err := s.cardRepository.Save(newCard)

	if err != nil {
		return nil, err
	}

	return cardModel.Marshal(), nil
}
