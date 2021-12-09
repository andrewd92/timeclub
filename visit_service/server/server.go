package server

import (
	"github.com/andrewd92/timeclub/visit_service/infrastructure/migration"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var (
	router = gin.Default()
)

func StartApplication() {
	initLogger()
	initConfig()

	migration.Run()

	mapUrls()

	log.WithField("port", viper.GetString("server.port.http")).Info("HTTP port overridden by config")
	log.WithField("port", os.Getenv("SERVER_PORT_HTTP")).Info("HTTP port overridden by env variable")

	port := viper.GetString("server.port.http")

	registerServiceWithConsul()

	err := router.Run(":" + port)
	if err != nil {
		log.WithError(err).Fatal("Run server error")
	}

	log.Infof("Server run on port :%s", port)
}
