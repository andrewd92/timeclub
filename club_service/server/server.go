package server

import (
	"fmt"
	"github.com/andrewd92/timeclub/club_service/api"
	"github.com/andrewd92/timeclub/club_service/grpc/club_server"
	"github.com/andrewd92/timeclub/club_service/infrastructure/migration"
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
	initConfig()
	initLogger()

	migration.Run()

	go runGrpcServer()
	runHttpServer()
}

func runHttpServer() {
	mapUrls()

	err := router.Run(":8080")
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

func initConfig() {
	viper.SetConfigName("local") // name of config file (without extension)
	viper.SetConfigType("yaml")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")    // path to look for the config file in
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./club_service/config")
	viper.AddConfigPath("$HOME/config") // call multiple times to add many search paths
	viper.AddConfigPath(".")            // optionally look for config in the working directory

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Fatal("Can not read config file")
	}
}

func initLogger() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
}
