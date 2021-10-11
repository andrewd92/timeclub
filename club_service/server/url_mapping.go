package server

import "github.com/gin-gonic/gin"

func mapUrls() {
	router.GET("/ping", pong)
}

func pong(c *gin.Context) {
	c.String(200, "pong")
}
