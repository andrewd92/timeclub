package test

import (
	"bytes"
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

func Test_Create(t *testing.T) {
	url := fmt.Sprintf("http://%s:%s/public/api/v1/", viper.GetString("server.host"), viper.GetString("server.port.http"))
	log.WithField("url", url).Debug("Test url")
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"club_id": 1,
		"card_id": 1,
	})
	requestBody := bytes.NewBuffer(postBody)

	response, err := client.Post(url, "application/json", requestBody)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	body, _ := io.ReadAll(response.Body)

	log.WithField("body", string(body)).Debug("create visit response")

	validateResponse(body, t)
}

func validateResponse(body []byte, t *testing.T) {
	var responseBody map[string]interface{}
	err := json.Unmarshal(body, &responseBody)
	assert.Nil(t, err)

	log.WithField("responseBody", responseBody).Debug("Response body")

	assert.True(t, responseBody["id"].(float64) > 0)
	start, err := time.Parse(utils.TimeFormat, responseBody["start"].(string))
	assert.Nil(t, err)
	hourBefore := time.Now().Add(-1 * time.Hour)
	assert.True(t, start.Unix() > hourBefore.Unix())
}
