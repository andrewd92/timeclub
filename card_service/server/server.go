package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	runHttpServer()
}

func runHttpServer() {
	mapUrls()

	err := router.Run(":8082")
	if err != nil {
		fmt.Println("Err: ", err.Error())
	}
}
