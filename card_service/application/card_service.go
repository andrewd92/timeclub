package application

import (
	cardPkg "github.com/andrewd92/timeclub/card_service/domain/card"
	"github.com/andrewd92/timeclub/card_service/domain/card/card_template"
	discountPkg "github.com/andrewd92/timeclub/card_service/domain/discount"
	"github.com/andrewd92/timeclub/card_service/infrastructure/repository/card_repository"
	"github.com/andrewd92/timeclub/card_service/infrastructure/repository/card_template_repository"
)

type cardServiceImpl struct {
	cardRepository         cardPkg.Repository
	cardTemplateRepository card_template.Repository
}

type CardService interface {
	Create(clubId int64, discount float32, name string) (interface{}, error)
	CreateTemplate(clubId int64, discount float32, name string) (interface{}, error)
	All() ([]interface{}, error)
	ById(id int64) (interface{}, error)
}

var service CardService

func CardServiceInstance() CardService {
	if nil == service {
		service = &cardServiceImpl{
			cardRepository:         card_repository.Instance(),
			cardTemplateRepository: card_template_repository.Instance(),
		}
	}

	return service
}

func (s cardServiceImpl) Create(clubId int64, discount float32, name string) (interface{}, error) {
	newCard := cardPkg.NewCard(discountPkg.Discount(discount), name, clubId)
	cardModel, err := s.cardRepository.Save(newCard)

	if err != nil {
		return nil, err
	}

	return cardModel.Marshal(), nil
}

func (s cardServiceImpl) CreateTemplate(clubId int64, discount float32, name string) (interface{}, error) {
	template := card_template.NewCardTemplate(discountPkg.Discount(discount), name, clubId)
	templateModel, err := s.cardTemplateRepository.Save(template)

	if err != nil {
		return nil, err
	}

	return templateModel.Marshal(), nil
}

func (s cardServiceImpl) All() ([]interface{}, error) {
	cards, err := s.cardRepository.All()
	if err != nil {
		return nil, err
	}

	return cardPkg.MarshalAll(cards), nil
}

func (s cardServiceImpl) ById(id int64) (interface{}, error) {
	card, err := s.cardRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return card.Marshal(), nil
}
