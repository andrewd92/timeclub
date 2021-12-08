package visit_controller

import "github.com/andrewd92/timeclub/visit_service/application/visit_service"

var instance *VisitController

func Instance() *VisitController {
	if nil == instance {
		instance = &VisitController{visitService: visit_service.Instance()}
	}

	return instance
}
