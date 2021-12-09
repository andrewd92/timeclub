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

	visitController := visit_controller.Instance()
	router.GET("/public/api/v1/:club_id", visitController.All)
	router.GET("/public/api/v1/:club_id/time/:time", visitController.ForTime)
	router.POST("/public/api/v1/", visitController.Create)

	router.GET("/public/api/v1/events", event_controller.All)
	router.POST("/public/api/v1/event", event_controller.Create)
}

func health(c *gin.Context) {
	log.Debug("Health endpoint called")

	c.String(http.StatusOK, "Up")
}
