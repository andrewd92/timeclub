package test

import (
	"encoding/json"
	"fmt"
	"github.com/andrewd92/timeclub/visit_service/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

func Test_GetForTime(t *testing.T) {
	now := time.Now().Add(time.Hour)

	url := fmt.Sprintf("http://%s:%s/visits/1/time/%s",
		viper.GetString("server.host"),
		viper.GetString("server.port.http"),
		now.Format(utils.TimeFormat),
	)
	log.WithField("url", url).Debug("Test url")

	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}

	response, err := client.Get(url)

	assert.Nil(t, err)
	assert.NotNil(t, response)

	body, _ := io.ReadAll(response.Body)

	log.Info(string(body))

	var responseBody []map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	assert.Nil(t, err)

	firstVisit := responseBody[0]

	log.WithField("first_visit", firstVisit).Info("First visit")

	assert.NotNil(t, firstVisit["start"])
	assert.Equal(t, float64(1), firstVisit["id"])
	assert.Equal(t, float64(1200), firstVisit["price"])
	assert.Equal(t, float64(120), firstVisit["duration"])
}
