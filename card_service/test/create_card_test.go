package test

import (
	"bytes"
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

func Test_Create(t *testing.T) {

	url := fmt.Sprintf("http://%s:%s/card", viper.GetString("server.host"), viper.GetString("server.port.http"))
	log.WithField("url", url).Debug("Test host URL")
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"club_id":  1,
		"name":     "Best Clients Card",
		"discount": 30.1,
	})
	requestBody := bytes.NewBuffer(postBody)

	response, err := client.Post(url, "application/json", requestBody)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	body, _ := io.ReadAll(response.Body)

	log.WithField("body", body).Info("Create card response body")

	expected := `{"id":4,"discount":30.1,"name":"Best Clients Card","club_id":1}`
	assert.Equal(t, expected, string(body))
}
