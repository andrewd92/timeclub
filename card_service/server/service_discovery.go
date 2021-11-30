package server

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	address := hostname()

	registration.ID = address
	registration.Name = viper.GetString("service.name")

	registration.Address = address
	registration.Port = httpPort()

	registration.TaggedAddresses = map[string]consulapi.ServiceAddress{
		"http": {Address: address, Port: httpPort()},
		"grpc": {Address: address, Port: grpcPort()},
	}

	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%d/health", address, httpPort())
	registration.Check.Interval = viper.GetString("consul.check.interval")
	registration.Check.Timeout = viper.GetString("consul.check.timeout")

	registrationErr := consul.Agent().ServiceRegister(registration)
	if registrationErr != nil {
		log.WithError(registrationErr).Error("Can not register in consul")
	}
}

func httpPort() int {
	port := viper.GetInt("server.port.http")
	if port == 0 {
		log.Error("Service http port not found in config")
		port = 8080
	}

	return port
}

func grpcPort() int {
	port := viper.GetInt("server.grpc.http")
	if port == 0 {
		log.Error("Service grpc port not found in config")
		port = 9084
	}

	return port
}

func hostname() string {
	hostname := viper.GetString("server.host")
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
