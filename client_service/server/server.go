package server

import (
	"fmt"
	"github.com/andrewd92/timeclub/client_service/api"
	"github.com/andrewd92/timeclub/client_service/grpc/client_service"
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

	err := router.Run(":8084")
	if err != nil {
		fmt.Println("Err: ", err.Error())
	}
}

func runGrpcServer() {
	listen, listenErr := net.Listen("tcp", ":9084")
	if listenErr != nil {
		log.Fatalf("failed to listen: %v", listenErr)
	}

	grpcServer := grpc.NewServer()
	clientServer, clientServerErr := client_service.Instance()

	if clientServerErr != nil {
		log.Fatal(clientServerErr)
	}

	api.RegisterClientServiceServer(grpcServer, clientServer)
	//run grpc on :9090

	log.Println("Listen and serve GRPC on :9084")
	err := grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
