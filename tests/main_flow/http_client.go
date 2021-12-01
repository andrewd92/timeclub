package main_flow

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"time"
)

func createHttpClient() *resty.Client {
	serverUrl := viper.GetString("server.url")

	client := resty.New()
	client.SetBaseURL(serverUrl)
	client.SetTimeout(1 * time.Second)

	return client
}
