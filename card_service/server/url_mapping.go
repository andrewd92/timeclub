package server

import (
	"github.com/andrewd92/timeclub/card_service/controller/card_controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func mapUrls() {
	router.GET("/health", health)

	router.GET("/card", card_controller.All)
	router.GET("/card/:id", card_controller.ById)
	router.POST("/card", card_controller.Create)
	router.GET("/card/templates", card_controller.Templates)
	router.POST("/card/template", card_controller.CreateTemplate)
}
func health(c *gin.Context) {
	log.Debug("Health endpoint called")

	c.String(http.StatusOK, "Up")
}
