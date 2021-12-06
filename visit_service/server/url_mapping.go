package server

import (
	"github.com/andrewd92/timeclub/visit_service/controller/event_controller"
	"github.com/andrewd92/timeclub/visit_service/controller/visit_controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func mapUrls() {
	router.GET("/health", health)

	router.GET("/visits/:club_id", visit_controller.All)
	router.POST("/visit", visit_controller.Create)

	router.GET("/events", event_controller.All)
	router.POST("/event", event_controller.Create)
}

func health(c *gin.Context) {
	log.Debug("Health endpoint called")

	c.String(http.StatusOK, "Up")
}
