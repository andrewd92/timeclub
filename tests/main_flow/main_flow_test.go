package main_flow

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_main_flow(t *testing.T) {
	client := createHttpClient()

	response, err := sendCreateClubRequest(client)

	assert.Nil(t, err)
	clubId := validateCreateClubResponse(t, response)

	log.WithField("club_id", clubId).Info("Club created")

	response, err = sendCreateCardRequest(client)

	assert.Nil(t, err)
	cardId := validateCreateCardResponse(t, response)

	log.WithField("card_id", cardId).Info("Card created")
}

func sendCreateClubRequest(client *resty.Client) (*resty.Response, error) {
	return client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"name":"Test Club","open_time":"12:00","price_list_id":[{"price_period":{"from":0,"to":60},"value_per_minute":5},{"price_period":{"from":60,"to":360},"value_per_minute":10}],"currency_id":1}`).
		SetResult(map[string]interface{}{}).
		Post("/club/api/v1/create")
}

func validateCreateClubResponse(t *testing.T, response *resty.Response) int64 {
	assert.Equal(t, http.StatusOK, response.StatusCode())

	body := *response.Result().(*map[string]interface{})
	id := body["id"]
	if nil == id {
		t.Fatal("Can not create club. No club id in response from server")
	}
	assert.Len(t, body["prices"], 2)

	return int64(id.(float64))
}

func sendCreateCardRequest(client *resty.Client) (*resty.Response, error) {
	return client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"name": "Best Clients Card", "club_id": 1, "discount": 15.4}`).
		SetResult(map[string]interface{}{}).
		Post("/card/api/v1/create")
}

func validateCreateCardResponse(t *testing.T, response *resty.Response) int64 {
	assert.Equal(t, http.StatusOK, response.StatusCode())

	body := *response.Result().(*map[string]interface{})
	id := body["id"]
	if nil == id {
		t.Fatal("Can not create card. No club id in response from server")
	}
	assert.Equal(t, body["discount"], 15.4)

	return int64(id.(float64))
}
