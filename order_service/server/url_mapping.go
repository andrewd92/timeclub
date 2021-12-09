package server

import (
	"github.com/andrewd92/timeclub/order_service/controller/order_controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func mapUrls() {
	orderController := order_controller.Instance()

	router.GET("/health", health)

	router.GET("/public/api/v1/:visit_id", orderController.ForVisit)
	router.POST("/public/api/v1/", orderController.Create)
	router.POST("/public/api/v1/pay", orderController.Pay)
	router.DELETE("/public/api/v1/:id", orderController.Cancel)
}

func health(c *gin.Context) {
	log.Debug("Health endpoint called")

	c.String(http.StatusOK, "Up")
}
