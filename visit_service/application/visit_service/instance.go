package visit_service

import (
	"github.com/andrewd92/timeclub/visit_service/client/club_service"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/repository/visit_repository"
)

var service VisitService

func Instance() VisitService {
	if nil == service {
		service = &visitServiceImpl{
			visitRepository:   visit_repository.Instance(),
			clubServiceClient: club_service.Instance(),
			visitMarshaller:   visit.MarshallerImpl{},
			visitFactory:      visit.FactoryImpl{},
		}
	}

	return service
}
