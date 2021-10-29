package card_controller

import (
	"github.com/andrewd92/timeclub/club_service/application"
	"github.com/andrewd92/timeclub/club_service/domain/card/card_template"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/card_template_repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Templates(c *gin.Context) {
	repository := card_template_repository.Instance()

	templates := repository.GetAll()

	response := card_template.MarshalAll(templates)

	c.JSON(http.StatusOK, response)
}

type createTemplateRequest struct {
	Discount float32 `json:"discount"`
	ClubId   int64   `json:"club_id"`
	Name     string  `json:"name"`
}

func CreateTemplate(c *gin.Context) {
	request := createTemplateRequest{}

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
