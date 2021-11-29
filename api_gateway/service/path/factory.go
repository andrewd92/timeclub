package path

import "github.com/andrewd92/timeclub/api_gateway/service/service_discovery"

var service Service

func Instance() Service {
	if service == nil {
		service = &serviceImpl{serviceDiscovery: service_discovery.Instance()}
	}

	return service
}
