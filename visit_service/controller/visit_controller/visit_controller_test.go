package visit_controller

import (
	"errors"
	visitServiceMock "github.com/andrewd92/timeclub/visit_service/application/visit_service/mocks"
	"github.com/andrewd92/timeclub/visit_service/utils/gin_test_utils"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func Test_ForTime(t *testing.T) {
	controller, service := newTestInstance()
	timeStr := "2222-03-18 23:01:06"
	clubId := int64(1)

	c, w := gin_test_utils.Init(gin_test_utils.Request{Params: map[string]string{
		clubIdKey: strconv.FormatInt(clubId, 10),
		"time":    timeStr,
	}})

	requestTime, _ := controller.parseTime(timeStr)

	responseMap := map[string]string{"response": "ok"}
	service.On("All", clubId, requestTime).Return([]interface{}{responseMap}, nil)

	controller.ForTime(c)

	expected := `[{"response":"ok"}]`
	response, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, string(response))
}

func Test_ForTime_RespondBadRequest_WhenWrongTimeFormat(t *testing.T) {
	controller, _ := newTestInstance()
	timeStr := "123"
	clubId := int64(1)

	c, w := gin_test_utils.Init(gin_test_utils.Request{Params: map[string]string{
		clubIdKey: strconv.FormatInt(clubId, 10),
		"time":    timeStr,
	}})

	controller.ForTime(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_ForTime_ResponseStatus500_WhenVisitServiceError(t *testing.T) {
	controller, service := newTestInstance()
	timeStr := "2222-03-18 23:01:06"
	clubId := int64(1)

	c, w := gin_test_utils.Init(gin_test_utils.Request{Params: map[string]string{
		clubIdKey: strconv.FormatInt(clubId, 10),
		"time":    timeStr,
	}})

	requestTime, _ := controller.parseTime(timeStr)

	service.On("All", clubId, requestTime).Return(nil, errors.New("test error"))

	controller.ForTime(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_Create(t *testing.T) {
	controller, service := newTestInstance()
	now := time.Now()
	timeNowMock = &now

	c, w := gin_test_utils.Init(gin_test_utils.Request{Body: `{"club_id":3, "card_id": 2}`})

	service.On("Create", int64(3), int64(2), *timeNowMock).
		Return(map[string]string{"response": "created"}, nil)

	controller.Create(c)

	expected := `{"response":"created"}`
	response, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, string(response))
}

func Test_Create_ResponseStatus400_WhenWrongRequest(t *testing.T) {
	controller, _ := newTestInstance()
	now := time.Now()
	timeNowMock = &now

	c, w := gin_test_utils.Init(gin_test_utils.Request{Body: `{"club_id":"x", "card_id": "y"}`})

	controller.Create(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_Create_ResponseStatus500_WhenServiceError(t *testing.T) {
	controller, service := newTestInstance()
	now := time.Now()
	timeNowMock = &now

	c, w := gin_test_utils.Init(gin_test_utils.Request{Body: `{"club_id":3, "card_id": 2}`})

	service.On("Create", int64(3), int64(2), *timeNowMock).
		Return(nil, errors.New("test error"))

	controller.Create(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func newTestInstance() (*VisitController, *visitServiceMock.VisitService) {
	service := new(visitServiceMock.VisitService)

	return &VisitController{visitService: service}, service
}
