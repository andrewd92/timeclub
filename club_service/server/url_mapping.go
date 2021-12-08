package server

import (
	"github.com/andrewd92/timeclub/club_service/controller/club_controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func mapUrls() {
	//v1
	router.GET("/health", health)
	router.GET("/public/api/v1/", club_controller.All)
	router.POST("/public/api/v1/", club_controller.Create)

}

func health(c *gin.Context) {
	log.Debug("Health endpoint called")

	c.String(http.StatusOK, "Up")
}
