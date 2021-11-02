package server

import (
	"fmt"
	"github.com/andrewd92/timeclub/club_service/api"
	"github.com/andrewd92/timeclub/club_service/grpc/club_server"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	router = gin.Default()
)

func StartApplication() {
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
		log.Fatal(clubServerErr)
	}

	api.RegisterClubServiceServer(grpcServer, clubServer)
	//run grpc on :9090

	log.Println("Listen and serve GRPC on :9080")
	err := grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
