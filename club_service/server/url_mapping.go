package server

import (
	"github.com/andrewd92/timeclub/club_service/controller/club_controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func mapUrls() {
	router.GET("/clubs", club_controller.All)
	router.GET("/health", health)
}

func health(c *gin.Context) {
	log.Debug("Health endpoint called")

	c.String(http.StatusOK, "Up")
}
