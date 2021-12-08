package visit_controller

import (
	"bytes"
	"errors"
	visitServiceMock "github.com/andrewd92/timeclub/visit_service/application/visit_service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func Test_ForTime(t *testing.T) {
	controller, service := newTestInstance()
	timeStr := "2222-03-18 23:01:06"
	clubId := int64(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key:   clubIdKey,
			Value: strconv.FormatInt(clubId, 10),
		},
		{
			Key:   "time",
			Value: timeStr,
		},
	}

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

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key:   clubIdKey,
			Value: strconv.FormatInt(clubId, 10),
		},
		{
			Key:   "time",
			Value: timeStr,
		},
	}

	controller.ForTime(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_ForTime_ResponseStatus500_WhenVisitServiceError(t *testing.T) {
	controller, service := newTestInstance()
	timeStr := "2222-03-18 23:01:06"
	clubId := int64(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key:   clubIdKey,
			Value: strconv.FormatInt(clubId, 10),
		},
		{
			Key:   "time",
			Value: timeStr,
		},
	}

	requestTime, _ := controller.parseTime(timeStr)

	service.On("All", clubId, requestTime).Return(nil, errors.New("test error"))

	controller.ForTime(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_Create(t *testing.T) {
	controller, service := newTestInstance()
	now := time.Now()
	timeNowMock = &now

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	requestStr := bytes.NewBuffer([]byte(`{"club_id":3, "card_id": 2}`))
	c.Request, _ = http.NewRequest(http.MethodPost, "/public/api/v1/", requestStr)

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

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	requestStr := bytes.NewBuffer([]byte(`{"club_id":"x", "card_id": "y"}`))
	c.Request, _ = http.NewRequest(http.MethodPost, "/public/api/v1/", requestStr)

	controller.Create(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_Create_ResponseStatus500_WhenServiceError(t *testing.T) {
	controller, service := newTestInstance()
	now := time.Now()
	timeNowMock = &now

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	requestStr := bytes.NewBuffer([]byte(`{"club_id":3, "card_id": 2}`))
	c.Request, _ = http.NewRequest(http.MethodPost, "/public/api/v1/", requestStr)

	service.On("Create", int64(3), int64(2), *timeNowMock).
		Return(nil, errors.New("test error"))

	controller.Create(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func newTestInstance() (*VisitController, *visitServiceMock.VisitService) {
	service := new(visitServiceMock.VisitService)

	return &VisitController{visitService: service}, service
}
