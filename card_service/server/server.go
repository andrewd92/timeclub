package server

import (
	"github.com/andrewd92/timeclub/card_service/infrastructure/migration"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	router = gin.Default()
)

func StartApplication() {
	initLogger()
	initConfig()

	migration.Run()

	registerServiceWithConsul()

	runHttpServer()
}

func runHttpServer() {
	mapUrls()

	port := viper.GetString("server.port.http")
	log.WithField("port", port).Info("HTTP port overridden by config")

	err := router.Run(":" + port)
	if err != nil {
		log.WithError(err).Fatal("can not run http server")
	}
}
