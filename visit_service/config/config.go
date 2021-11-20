package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"path/filepath"
)

type Config struct {
	Server struct {
		Port struct {
			Http string `yaml:"http"`
			Grpc string `yaml:"grpc"`
		} `yaml:"port"`
	} `yaml:"server"`
	Client struct {
		Club struct {
			Grpc struct {
				Url string `yaml:"url"`
			} `yaml:"grpc"`
		} `yaml:"club"`
	} `yaml:"client"`
}

var config *Config
var configFilePath = "/visit_service/config.yml"

func Instance() Config {
	if config != nil {
		return *config
	}

	filename, executableErr := filepath.Abs("./")
	if executableErr != nil {
		panic(executableErr)
	}
	f, err := os.Open(path.Join(path.Dir(filename), configFilePath))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	return *config
}

func Init() {
	viper.SetConfigName("config")          // name of config file (without extension)
	viper.SetConfigType("yaml")            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")              // path to look for the config file in
	viper.AddConfigPath("./visit_service") // path to look for the config file in
	viper.AddConfigPath("$HOME/config")    // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory

	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
