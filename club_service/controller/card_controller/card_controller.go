package card_controller

import (
	"github.com/andrewd92/timeclub/club_service/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createCardRequest struct {
	Discount float32 `json:"discount"`
	ClubId   int64   `json:"club_id"`
	Name     string  `json:"name"`
}

func Create(c *gin.Context) {
	request := createCardRequest{}

	bindingErr := c.BindJSON(&request)

	if bindingErr != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	service := application.CardServiceInstance()

	response, serviceErr := service.CreateTemplate(request.ClubId, request.Discount, request.Name)

	if serviceErr != nil {
		c.String(http.StatusInternalServerError, serviceErr.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
