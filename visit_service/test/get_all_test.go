package test

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestGetAll(t *testing.T) {

	url := fmt.Sprintf("http://%s:%s/visits", viper.GetString("server.host"), viper.GetString("server.port.http"))
	fmt.Println("Test URL: " + url)
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
	assert.NotNil(t, responseBody[0]["start"])

	expected := `[{"id":1,"start":"` + responseBody[0]["start"].(string) + `","client_name":"Sasha","club_id":1,"comment":"","price":540,"currency":"$","duration":60,"card_id":1}]`
	assert.Equal(t, expected, string(body))
}
