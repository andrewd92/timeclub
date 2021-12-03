package card_controller

import (
	"github.com/andrewd92/timeclub/card_service/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	response, serviceErr := service.Create(request.ClubId, request.Discount, request.Name)

	if serviceErr != nil {
		c.String(http.StatusInternalServerError, serviceErr.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func All(c *gin.Context) {
	service := application.CardServiceInstance()

	response, serviceErr := service.All()

	if serviceErr != nil {
		c.String(http.StatusInternalServerError, serviceErr.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func ById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	service := application.CardServiceInstance()

	response, serviceErr := service.ById(id)

	if serviceErr != nil {
		c.String(http.StatusInternalServerError, serviceErr.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
