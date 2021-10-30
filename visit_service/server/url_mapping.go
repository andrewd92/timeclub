package server

import (
	"github.com/andrewd92/timeclub/visit_service/controller/event_controller"
	"github.com/andrewd92/timeclub/visit_service/controller/visit_controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mapUrls() {
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	router.GET("/visits", visit_controller.All)
	router.POST("/visit", visit_controller.Create)

	router.GET("/events", event_controller.All)
	router.POST("/event", event_controller.Create)
}
