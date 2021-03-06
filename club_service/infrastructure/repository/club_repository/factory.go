package club_repository

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	"github.com/andrewd92/timeclub/club_service/infrastructure/dao/club_dao"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/price_list_repository"
)

var repository club.Repository

func Instance() club.Repository {
	if nil != repository {
		return repository
	}

	repository = &ClubDBRepository{
		dao:                 club_dao.NewClubDao(connection.Instance()),
		priceListRepository: price_list_repository.Instance(),
	}

	return repository
}

func MockTestRepository(mock club.Repository) {
	repository = mock
}
