package card_repository

import (
	"github.com/andrewd92/timeclub/card_service/domain/card"
	discountPkg "github.com/andrewd92/timeclub/card_service/domain/discount"
	"github.com/andrewd92/timeclub/card_service/infrastructure/dao/card_dao"
)

type CardDBRepository struct {
	dao card_dao.CardDao
}

func (c CardDBRepository) GetById(id int64) (*card.Card, error) {
	model, err := c.dao.GetById(id)
	if err != nil {
		return nil, err
	}

	return modelToEntity(model), nil
}

func (c CardDBRepository) All() ([]*card.Card, error) {
	models, err := c.dao.GetAll()

	if err != nil {
		return nil, err
	}

	cards := make([]*card.Card, len(models), len(models))

	for i, model := range models {
		cards[i] = modelToEntity(&model)
	}

	return cards, nil
}

func (c CardDBRepository) Save(card *card.Card) (*card.Card, error) {
	id, err := c.dao.Insert(card)
	if err != nil {
		return nil, err
	}

	return card.WithId(id), nil
}

func modelToEntity(model *card_dao.CardModel) *card.Card {
	clubId := int64(0)
	if model.ClubId.Valid {
		clubId = model.ClubId.Int64
	}
	return card.NewCard(discountPkg.Discount(model.Discount), model.Name, clubId).WithId(model.Id)
}
