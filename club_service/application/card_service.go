package application

import (
	"github.com/andrewd92/timeclub/club_service/domain/card"
	"github.com/andrewd92/timeclub/club_service/domain/card/card_template"
	discountPkg "github.com/andrewd92/timeclub/club_service/domain/discount"
	visitPkg "github.com/andrewd92/timeclub/club_service/domain/visit"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/card_repository"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/card_template_repository"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/visit_repository"
	"sort"
)

type cardServiceImpl struct {
	cardRepository         card.Repository
	cardTemplateRepository card_template.Repository
	visitRepository        visitPkg.Repository
}

type CardService interface {
	Create(cafeId int64, discount float32, name string) (interface{}, error)
	CreateTemplate(clubId int64, discount float32, name string) (interface{}, error)
	GetMinAvailableCard(clubId int64) (int64, error)
}

var service CardService

func CardServiceInstance() CardService {
	if nil == service {
		service = &cardServiceImpl{
			cardRepository:         card_repository.Instance(),
			cardTemplateRepository: card_template_repository.Instance(),
			visitRepository:        visit_repository.Instance(),
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

func (s cardServiceImpl) CreateTemplate(clubId int64, discount float32, name string) (interface{}, error) {
	template := card_template.NewCardTemplate(discountPkg.NewDiscount(discount), name, clubId)
	templateModel, err := s.cardTemplateRepository.Save(template)

	if err != nil {
		return nil, err
	}

	return templateModel.Marshal(), nil
}

func (s cardServiceImpl) GetMinAvailableCard(clubId int64) (int64, error) {
	cardIds := make([]int64, 0)

	for _, visit := range s.visitRepository.GetAll() {
		if visit.Club().Id() == clubId {
			cardIds = append(cardIds, visit.Card().Id())
		}
	}

	sort.Slice(cardIds, func(i, j int) bool {
		return cardIds[i] < cardIds[j]
	})

	result := int64(0)

	for i, id := range cardIds {
		if len(cardIds) > i+1 && cardIds[i+1]-id > 1 {
			result = id + 1
			break
		}
	}

	if 0 == result {
		if 0 == len(cardIds) {
			result = 1
		} else {
			result = cardIds[len(cardIds)-1] + 1
		}
	}

	return result, nil
}
