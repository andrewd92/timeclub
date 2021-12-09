package order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
}

func Instance() *OrderController {
	return &OrderController{}
}

func (oc OrderController) ForVisit(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (oc OrderController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     1,
		"visits": []int{1, 2, 3},
	})
}

func (oc OrderController) Pay(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (oc OrderController) Cancel(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
