package proxy

var service Service

func Instance() Service {
	if nil == service {
		service = &serviceImpl{}
	}

	return service
}
