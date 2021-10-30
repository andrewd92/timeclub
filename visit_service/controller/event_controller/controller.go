package event_controller

import (
	"github.com/andrewd92/timeclub/visit_service/application/event_service"
	"github.com/andrewd92/timeclub/visit_service/domain/event"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/repository/event_repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func All(c *gin.Context) {
	eventRepository := event_repository.Instance()

	events := eventRepository.GetAll()

	c.JSON(http.StatusOK, event.MarshalAll(events))
}

type createEventRequest struct {
	Name     string  `json:"name"`
	Tag      string  `json:"tag"`
	Discount float32 `json:"discount"`
	Start    string  `json:"start"`
	End      string  `json:"end"`
}

func Create(c *gin.Context) {
	request := createEventRequest{}

	bindingErr := c.BindJSON(&request)

	if bindingErr != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	service := event_service.Instance()

	response, serviceErr := service.Create(request.Name, request.Tag, request.Discount, request.Start, request.End)

	if serviceErr != nil {
		c.String(http.StatusInternalServerError, serviceErr.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
