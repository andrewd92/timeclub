package test

import (
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

	url := fmt.Sprintf("http://%s:%s/card", viper.GetString("server.host"), viper.GetString("server.port.http"))
	log.WithField("url", url).Debug("Test host URL")
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	response, err := client.Get(url)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	body, _ := io.ReadAll(response.Body)

	log.WithField("body", body).Debug("Response body")

	expected := `[{"id":0,"discount":0,"name":"Guest Card","club_id":null},{"id":0,"discount":0,"name":"Guest Card","club_id":null},{"id":0,"discount":0,"name":"Guest Card","club_id":null}]`
	assert.Equal(t, expected, string(body))
}
