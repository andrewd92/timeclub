package service_discovery

var service Service

func Instance() Service {
	if nil == service {
		service = &serviceImpl{}
	}

	return service
}
