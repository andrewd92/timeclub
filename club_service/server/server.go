package server

import (
	"fmt"
	"github.com/andrewd92/timeclub/club_service/api"
	"github.com/andrewd92/timeclub/club_service/grpc/club_server"
	"github.com/andrewd92/timeclub/club_service/infrastructure/migration"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

var (
	router = gin.Default()
)

func StartApplication() {
	initLogger()
	initConfig()

	migration.Run()

	registerServiceWithConsul()

	go runGrpcServer()
	runHttpServer()
}

func runHttpServer() {
	mapUrls()

	err := router.Run(":" + strconv.Itoa(port()))
	if err != nil {
		fmt.Println("Err: ", err.Error())
	}
}

func runGrpcServer() {
	listen, listenErr := net.Listen("tcp", ":9080")
	if listenErr != nil {
		log.Fatalf("failed to listen: %v", listenErr)
	}

	grpcServer := grpc.NewServer()
	clubServer, clubServerErr := club_server.Instance()

	if clubServerErr != nil {
		log.WithError(clubServerErr).Fatal("Can not instantiate club GRPC server")
	}

	api.RegisterClubServiceServer(grpcServer, clubServer)
	//run grpc on :9090

	log.Info("Listen and serve GRPC on :9080")
	err := grpcServer.Serve(listen)
	if err != nil {
		log.WithError(err).Fatal("GRPC Server ERROR")
	}
}
