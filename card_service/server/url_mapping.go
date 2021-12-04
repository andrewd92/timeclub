package server

import (
	"github.com/andrewd92/timeclub/card_service/controller/card_controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func mapUrls() {
	router.GET("/health", health)

	router.POST("/api/v1/create", card_controller.Create)
	router.GET("/api/v1/all", card_controller.All)
	router.GET("/api/v1/:id", card_controller.ById)
	router.GET("/card/templates", card_controller.Templates)
	router.POST("/card/template", card_controller.CreateTemplate)
}
func health(c *gin.Context) {
	log.Debug("Health endpoint called")

	c.String(http.StatusOK, "Up")
}
