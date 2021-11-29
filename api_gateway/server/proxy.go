package server

import (
	pathService "github.com/andrewd92/timeclub/api_gateway/service/path"
	"github.com/andrewd92/timeclub/api_gateway/service/proxy"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

func ReverseProxy(ctx *gin.Context) {
	path := ctx.Param("proxyPath")
	target, err := pathService.Instance().Target(path)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	if targetUrl, err := url.Parse(target); err != nil {
		log.WithError(err).
			WithField("target_url", targetUrl).
			WithField("path", path).
			Error("can not parse target url")
		ctx.AbortWithStatus(http.StatusInternalServerError)
	} else {
		proxy.Instance().Proxy(targetUrl).ServeHTTP(ctx.Writer, ctx.Request)
	}
}
