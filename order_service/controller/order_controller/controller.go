package order_controller

import (
	"github.com/andrewd92/timeclub/order_service/api/http/create_order"
	"github.com/andrewd92/timeclub/order_service/application/order_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	orderService order_service.OrderService
}

func Instance() *OrderController {
	return &OrderController{}
}

func (oc OrderController) ForVisit(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (oc OrderController) Create(c *gin.Context) {
	request := create_order.Request{}

	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	response, err := oc.orderService.Create(request.Visits)
	if err != nil {

	}

	c.JSON(http.StatusOK, response)
}

func (oc OrderController) Pay(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (oc OrderController) Cancel(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
