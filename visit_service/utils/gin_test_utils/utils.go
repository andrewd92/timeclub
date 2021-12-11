package gin_test_utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

type Request struct {
	Params map[string]string
	Body   string
}

func Init(request Request) (*gin.Context, *httptest.ResponseRecorder) {
	context, responseRecorder := initGin()

	if nil != request.Params {
		var requestParams []gin.Param

		for k, v := range request.Params {
			requestParams = append(requestParams, gin.Param{Key: k, Value: v})
		}

		context.Params = requestParams
	}

	if "" != request.Body {
		requestStrBuffer := bytes.NewBuffer([]byte(request.Body))
		context.Request, _ = http.NewRequest(http.MethodPost, "/", requestStrBuffer)
	}

	return context, responseRecorder
}

func initGin() (*gin.Context, *httptest.ResponseRecorder) {
	responseRecorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseRecorder)

	return context, responseRecorder
}
