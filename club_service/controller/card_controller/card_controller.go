package card_controller

import (
	"github.com/andrewd92/timeclub/club_service/application"
	"github.com/andrewd92/timeclub/club_service/domain/card/card_template"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/card_template_repository"
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

func MinAvailable(c *gin.Context) {
	clubId, parseErr := strconv.ParseInt(c.Param("clubId"), 10, 64)
	if parseErr != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	service := application.CardServiceInstance()

	minAvailableCard, serviceErr := service.GetMinAvailableCard(clubId)

	if serviceErr != nil {
		c.String(http.StatusInternalServerError, serviceErr.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"card_id": minAvailableCard})
}

func Templates(c *gin.Context) {
	repository := card_template_repository.Instance()

	templates := repository.GetAll()

	response := card_template.MarshalAll(templates)

	c.JSON(http.StatusOK, response)
}
