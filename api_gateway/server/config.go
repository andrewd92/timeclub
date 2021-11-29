package server

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func initConfig() {
	configName := os.Getenv("VIPER_CONFIG_NAME")
	if configName == "" {
		log.Warning("Can not find VIPER_CONFIG_NAME environment variable. Local config used")
		configName = "local"
	}

	viper.SetConfigName(configName) // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/")        // path to look for the config file in
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./api_gateway/config")
	viper.AddConfigPath("$HOME/config") // call multiple times to add many search paths
	viper.AddConfigPath(".")            // optionally look for config in the working directory

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Fatal("Can not read config file")
	}

	log.
		WithField("server_name", viper.GetString("service.name")).
		WithField("config_name", configName).
		Info("Viper config is ready")
}
