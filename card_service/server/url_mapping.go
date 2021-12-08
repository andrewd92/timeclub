package server

import (
	"github.com/andrewd92/timeclub/card_service/controller/card_controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func mapUrls() {
	router.GET("/health", health)

	router.POST("/public/api/v1/", card_controller.Create)
	router.GET("/public/api/v1/", card_controller.All)
	router.GET("/public/api/v1/:id", card_controller.ById)
	router.GET("/public/api/v1/templates", card_controller.Templates)
	router.POST("/public/api/v1/template", card_controller.CreateTemplate)
}
func health(c *gin.Context) {
	log.Debug("Health endpoint called")

	c.String(http.StatusOK, "Up")
}
