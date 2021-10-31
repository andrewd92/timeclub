package visit_service

import "github.com/andrewd92/timeclub/visit_service/infrastructure/repository/visit_repository"

var service VisitService

func Instance() VisitService {
	if nil == service {
		service = &visitServiceImpl{
			visitRepository: visit_repository.Instance(),
		}
	}

	return service
}
