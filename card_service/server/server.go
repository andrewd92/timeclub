package server

import (
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
