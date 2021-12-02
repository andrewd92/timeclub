package card_repository

import (
	"github.com/andrewd92/timeclub/card_service/domain/card"
	"github.com/andrewd92/timeclub/card_service/infrastructure/connection"
	"github.com/andrewd92/timeclub/card_service/infrastructure/dao/card_dao"
)

var repository card.Repository

func Instance() card.Repository {
	if nil != repository {
		return repository
	}

	repository = &CardDBRepository{
		dao: card_dao.NewCardSqlDao(connection.Instance()),
	}

	return repository
}
