package server

import (
	"github.com/andrewd92/timeclub/visit_service/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	config.Init()

	port := viper.GetString("server.port.http")

	log.WithField("port", viper.GetString("SERVER_PORT_HTTP")).Info("HTTP port overridden by config")
	log.WithField("port", os.Getenv("SERVER_PORT_HTTP")).Info("HTTP port overridden by env variable")

	err := router.Run(":" + port)
	if err != nil {
		log.WithError(err).Fatal("Run server error")
	}

	log.Infof("Server run on port :%s", port)
}
