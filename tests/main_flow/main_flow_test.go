package main_flow

import (
	"fmt"
	"github.com/andrewd92/timeclub/tests/utils"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

var clubId int64
var cardId int64
var visitId int64

func Test_main_flow(t *testing.T) {
	client := createHttpClient()

	response, err := sendCreateClubRequest(client)

	assert.Nil(t, err)
	clubId = validateCreateClubResponse(t, response)

	log.WithField("club_id", clubId).Info("Club created")

	response, err = sendCreateCardRequest(client)
	assert.Nil(t, err)

	cardId = validateCreateCardResponse(t, response)

	log.WithField("card_id", cardId).Info("Card created")

	response, err = sendCreateVisitRequest(client)
	assert.Nil(t, err)

	visitId = validateCreateVisitResponse(t, response)
	log.WithField("visit_id", visitId).Info("Visit created")

	response, err = sendGetVisits(client)
	assert.Nil(t, err)

	validateVisit(t, response)
}

func sendGetVisits(client *resty.Client) (*resty.Response, error) {
	now := time.Now().Add(time.Hour)

	url := fmt.Sprintf("/visit/visits/%d/time/%s", clubId, now.Format(utils.TimeFormat))

	log.WithField("url", url).Info("Get visits for time")

	return client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(make([]map[string]interface{}, 0)).
		Get(url)
}

func validateVisit(t *testing.T, response *resty.Response) {
	assert.Equal(t, http.StatusOK, response.StatusCode())

	log.Debug(string(response.Body()))

	body := *response.Result().(*[]map[string]interface{})

	log.WithField("body", body).Debug("Parsed body")

	found := false
	for _, visit := range body {
		if utils.ParseInt64(visit["id"]) != visitId {
			continue
		}

		found = true

		assert.Equal(t, visit["price"].(float64), float64(300))
		assert.Equal(t, visit["duration"].(float64), float64(60))
		assert.Equal(t, visit["currency"].(string), "$")
		assert.Equal(t, utils.ParseInt64(visit["club_id"]), clubId)
		assert.Equal(t, utils.ParseInt64(visit["card_id"]), cardId)
	}

	assert.True(t, found, "Created visit not stored")
}

func sendCreateVisitRequest(client *resty.Client) (*resty.Response, error) {
	createVisitRequest := map[string]int64{
		"club_id": clubId,
		"card_id": cardId,
	}
	return client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(createVisitRequest).
		SetResult(map[string]interface{}{}).
		Post("/visit/visit")
}

func validateCreateVisitResponse(t *testing.T, response *resty.Response) int64 {
	assert.Equal(t, http.StatusOK, response.StatusCode())

	body := *response.Result().(*map[string]interface{})
	id := body["id"]
	if nil == id {
		t.Fatal("Can not create visit. No visit id in response from server")
	}

	log.WithField("visit_body", body).Debug("Visit body")

	assert.Equal(t, cardId, utils.ParseInt64(body["card_id"]))
	assert.Equal(t, clubId, utils.ParseInt64(body["club_id"]))

	return utils.ParseInt64(id)
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
