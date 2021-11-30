package server

import (
	"github.com/andrewd92/timeclub/client_service/api"
	"github.com/andrewd92/timeclub/client_service/grpc/client_service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

var (
	router = gin.Default()
)

func StartApplication() {
	initLogger()
	initConfig()

	registerServiceWithConsul()

	go runGrpcServer()
	runHttpServer()
}

func runHttpServer() {
	port := viper.GetString("server.port.http")
	log.WithField("port", port).Info("HTTP port overridden by config")

	mapUrls()

	err := router.Run(":" + port)
	if err != nil {
		log.WithError(err).Fatal("can not run http server")
	}
}

func runGrpcServer() {
	port := viper.GetString("server.port.grpc")

	if port == "" {
		log.Error("can not find grpc port in config. User default: 9084")
		port = "9084"
	}

	listen, listenErr := net.Listen("tcp", ":"+port)
	if listenErr != nil {
		log.WithError(listenErr).WithField("port", port).Fatal("failed to listen grpc port")
	}

	grpcServer := grpc.NewServer()
	clientsGrpcService, clientServiceErr := client_service.Instance()

	if clientServiceErr != nil {
		log.WithError(clientServiceErr).Error("can not instantiate clients grpc service impl")
	}

	api.RegisterClientServiceServer(grpcServer, clientsGrpcService)

	log.WithField("port", port).Info("Listen and serve GRPC")
	err := grpcServer.Serve(listen)
	if err != nil {
		log.WithError(err).Fatal("can not run grpc server")
	}
}
