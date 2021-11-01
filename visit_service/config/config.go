package config

import (
	"gopkg.in/yaml.v2"
	"os"
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
var configFilePath string = "./config.yml"

func Instance() Config {
	if config != nil {
		return *config
	}

	f, err := os.Open(configFilePath)
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
