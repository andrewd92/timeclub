package path

import (
	"errors"
	"fmt"
	"github.com/andrewd92/timeclub/api_gateway/service/service_discovery"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Service interface {
	Target(path string) (string, error)
}

type serviceImpl struct {
	serviceDiscovery service_discovery.Service
}

func (s serviceImpl) Target(path string) (string, error) {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) <= 1 {
		log.WithField("path", path).Error("failed to parse target host from path")
		return "", errors.New("not found")
	}
	targetHost, err := s.serviceDiscovery.LookUpUrl(parts[0])
	if err != nil {
		log.WithError(err).Error("can not find service host")
		return "", errors.New("not found")
	}

	targetAddr := fmt.Sprintf(
		"%s/%s",
		targetHost, strings.Join(parts[1:], "/"),
	)

	log.WithField("path", path).
		WithField("targetPath", targetAddr).
		Debug("Revers proxy result")
	return targetAddr, nil
}
