package visit_repository

import (
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/dao/visit_dao"
)

var repository visit.Repository

func Instance() visit.Repository {
	if nil != repository {
		return repository
	}

	repository = &VisitDbRepository{dao: visit_dao.Instance()}

	return repository
}
