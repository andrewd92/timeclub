package order_controller

import (
	"bytes"
	"github.com/andrewd92/timeclub/order_service/application/order_service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Create(t *testing.T) {
	controller, orderService := testInstance()

	responseWriter := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseWriter)
	requestBytes := bytes.NewBuffer([]byte(`{"visits":[1,2,5]}`))
	c.Request, _ = http.NewRequest(http.MethodPost, "/public/api/v1/", requestBytes)

	orderService.On("Create", []int64{1, 2, 5}).Return(map[string]string{"response": "ok"}, nil)

	controller.Create(c)

	response, _ := io.ReadAll(responseWriter.Body)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, `{"response":"ok"}`, string(response))
}

func testInstance() (*OrderController, *mocks.OrderService) {
	orderService := new(mocks.OrderService)
	return &OrderController{orderService: orderService}, orderService
}
