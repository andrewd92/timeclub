package order_controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Create(t *testing.T) {
	controller := testInstance()

	responseWriter := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseWriter)
	requestBytes := bytes.NewBuffer([]byte(`{"visits":[1,2,3]}`))
	c.Request, _ = http.NewRequest(http.MethodPost, "/public/api/v1/", requestBytes)

	controller.Create(c)

	response, _ := io.ReadAll(responseWriter.Body)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, `{"id":1,"visits":[1,2,3]}`, string(response))
}

func testInstance() *OrderController {
	return &OrderController{}
}
