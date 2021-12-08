package test

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

func Test_GetAll(t *testing.T) {

	url := fmt.Sprintf("http://%s:%s/public/api/v1/", viper.GetString("server.host"), viper.GetString("server.port.http"))
	log.WithField("url", url).Debug("Test host URL")
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	response, err := client.Get(url)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	body, _ := io.ReadAll(response.Body)

	log.WithField("body", body).Debug("Response body")

	var responseBody []map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	assert.Nil(t, err)

	assert.True(t, len(responseBody) >= 3)
}
