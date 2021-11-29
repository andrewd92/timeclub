package service_discovery

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

type Service interface {
	LookUpUrl(serviceName string) (string, error)
}

type serviceImpl struct {
}

func (s serviceImpl) LookUpUrl(serviceName string) (string, error) {
	config := consulapi.DefaultConfig()
	config.Address = consulUrl()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.WithError(err).Error("Can not create consul client")
		return "", err
	}

	services, err := consul.Agent().Services()
	if err != nil {
		log.WithError(err).Error("Can not load services from consul")
		return "", err
	}

	match := make([]*consulapi.AgentService, 0)

	for _, service := range services {
		if service.Service == serviceName {
			match = append(match, service)
		}
	}

	if len(match) <= 0 {
		log.WithField("service", serviceName).Error("Can not find service in consul")
		return "", errors.New("not found")
	}

	rand.Seed(time.Now().Unix())
	service := match[rand.Intn(len(match))]
	address := service.Address
	port := service.Port

	url := fmt.Sprintf("http://%s:%v", address, port)

	log.WithField("service", serviceName).WithField("url", url).Debug("Service url found")

	return url, nil
}

func consulUrl() string {
	host := viper.GetString("consul.host")
	port := viper.GetString("consul.port")

	log.WithField("host", host).WithField("port", port).Debug("Consul config")
	return fmt.Sprintf("%s:%s", host, port)
}
