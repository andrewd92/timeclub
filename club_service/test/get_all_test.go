package test

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestGetAll(t *testing.T) {

	url := fmt.Sprintf("http://%s:%s/api/v1/clubs", viper.GetString("service.host"), viper.GetString("service.port"))
	fmt.Println("Test URL: " + url)
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	response, err := client.Get(url)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	body, _ := io.ReadAll(response.Body)

	expected := `[{"id":1,"name":"t1","open_time":"12:00","currency":"USD","prices":[{"price_period":{"from":0,"to":360},"value_per_minute":10}]}]`
	assert.Equal(t, expected, string(body))
}
