package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func InitConfig() {
	configName := "config"

	viper.SetConfigName(configName) // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./tests/config")
	viper.AddConfigPath("$HOME/config") // call multiple times to add many search paths
	viper.AddConfigPath(".")            // optionally look for config in the working directory

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Fatal("Can not read config file")
	}

	log.
		WithField("server_url", viper.GetString("server.url")).
		WithField("config_name", configName).
		Info("Viper config is ready")
}
