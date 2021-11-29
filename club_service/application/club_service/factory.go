package club_service

import "github.com/andrewd92/timeclub/club_service/infrastructure/repository/club_repository"

var service ClubService

func Instance() ClubService {
	if nil == service {
		service = &clubServiceImpl{repository: club_repository.Instance()}
	}

	return service
}
