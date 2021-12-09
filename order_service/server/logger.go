package server

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func initLogger() {
	logLevel := os.Getenv("APP_SERVER_LOG_LEVEL")
	if 0 == len(logLevel) {
		logLevel = "info"
	}

	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		FullTimestamp:          true,
		PadLevelText:           true,
	})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithField("level", logLevel).Warning("Can not parse log level. Info level set as default")
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}

}
