package server

import (
	"fmt"
	"github.com/andrewd92/timeclub/visit_service/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	config.Init()

	port := viper.GetString("server.port.http")

	log.Printf("ENV port :%s", viper.GetString("SERVER_PORT_HTTP"))
	fmt.Println("ENV: ", os.Getenv("SERVER_PORT_HTTP"))

	err := router.Run(":" + port)
	if err != nil {
		fmt.Println("Err: ", err.Error())
	}

	log.Printf("Server run on port :%s", port)
}
