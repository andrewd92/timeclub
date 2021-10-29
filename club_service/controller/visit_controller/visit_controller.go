package visit_controller

import (
	"github.com/andrewd92/timeclub/club_service/domain/visit"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/visit_repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func All(c *gin.Context) {
	repository := visit_repository.Instance()

	visits := repository.GetAll()

	response, responseErr := visit.MarshalAll(visits)

	if nil != responseErr {
		c.String(http.StatusInternalServerError, "All visits response building error")
		return
	}

	c.JSON(http.StatusOK, response)
}

type createRequest struct {
	ClubId int64 `json:"club_id"`
	CardId int64 `json:"card_id"`
}

func Create(c *gin.Context) {
	repository := visit_repository.Instance()

	request := createRequest{}

	bindingErr := c.BindJSON(&request)

	if bindingErr != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	createdVisit, createVisitErr := repository.CreateVisit(request.ClubId, request.CardId)

	if createVisitErr != nil {
		c.String(http.StatusBadRequest, createVisitErr.Error())
		return
	}

	response, responseErr := createdVisit.Marshal(time.Now())

	if responseErr != nil {
		c.String(http.StatusInternalServerError, "Response building error")
		return
	}

	c.JSON(http.StatusOK, response)
}
