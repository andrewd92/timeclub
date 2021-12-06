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

func TestGetAll(t *testing.T) {

	url := fmt.Sprintf("http://%s:%s/visits/1", viper.GetString("server.host"), viper.GetString("server.port.http"))
	log.WithField("url", url).Debug("Test url")
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	response, err := client.Get(url)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	body, _ := io.ReadAll(response.Body)

	var responseBody []map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	assert.Nil(t, err)

	firstVisit := responseBody[0]
	assert.NotNil(t, firstVisit["start"])
	assert.Equal(t, float64(1), firstVisit["id"])
}
