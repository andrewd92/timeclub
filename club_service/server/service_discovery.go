package server

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

func registerServiceWithConsul() {
	if viper.GetBool("consul.enabled") != true {
		return
	}

	config := consulapi.DefaultConfig()
	config.Address = consulUrl()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.WithError(err).Fatal("Can not create consul client")
	}

	var registration = new(consulapi.AgentServiceRegistration)

	registration.ID = viper.GetString("service.name")
	registration.Name = viper.GetString("service.name")

	address := hostname()
	registration.Address = address
	registration.Port = port()

	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%d/health", address, port())
	registration.Check.Interval = "20s"
	registration.Check.Timeout = "3s"

	registrationErr := consul.Agent().ServiceRegister(registration)
	if registrationErr != nil {
		log.WithError(registrationErr).Error("Can not register in consul")
	}
}

func port() int {
	port := viper.GetString("service.port")
	if len(port) == 0 {
		log.Error("Service port not found in config")
		port = "8080"
	}

	result, err := strconv.Atoi(port)
	if err != nil {
		log.WithError(err).WithField("port", port).Error("Can not parse port from config")
	}

	return result
}

func hostname() string {
	hostname := viper.GetString("service.host")
	if len(hostname) == 0 {
		log.Error("Service host not found in config")
		hostname = "localhost"
	}

	return hostname
}

func consulUrl() string {
	host := viper.GetString("consul.host")
	port := viper.GetString("consul.port")

	log.WithField("host", host).WithField("port", port).Debug("Consul config")
	return fmt.Sprintf("%s:%s", host, port)
}
