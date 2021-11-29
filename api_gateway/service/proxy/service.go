package proxy

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Service interface {
	Proxy(address *url.URL) *httputil.ReverseProxy
}

type serviceImpl struct {
}

func (s serviceImpl) Proxy(address *url.URL) *httputil.ReverseProxy {
	p := httputil.NewSingleHostReverseProxy(address)
	p.Director = func(request *http.Request) {
		request.Host = address.Host
		request.URL.Scheme = address.Scheme
		request.URL.Host = address.Host
		request.URL.Path = address.Path
	}

	p.ModifyResponse = func(response *http.Response) error {
		if response.StatusCode == http.StatusInternalServerError {
			s := readBody(response)

			logProxyError(address, response.StatusCode, s)

			response.Body = ioutil.NopCloser(bytes.NewReader([]byte("Internal server error")))
		} else if response.StatusCode > 300 {
			s := readBody(response)

			logProxyError(address, response.StatusCode, s)

			response.Body = ioutil.NopCloser(bytes.NewReader([]byte(s)))
		}

		return nil
	}

	return p
}

func logProxyError(address *url.URL, statusCode int, response string) {
	log.WithField("address", address).
		WithField("status_code", statusCode).
		WithField("response", response).
		Error("gateway proxy error")
}

func readBody(response *http.Response) string {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.WithError(err).Error("can not close response body buffer")
		}
	}(response.Body)

	all, _ := ioutil.ReadAll(response.Body)

	var body string
	if len(all) > 0 {
		body = string(all)
	}
	return body
}
