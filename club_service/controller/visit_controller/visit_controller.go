package visit_controller

import (
	"fmt"
	"github.com/andrewd92/timeclub/club_service/domain/visit"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/visit_repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func All(c *gin.Context) {
	repository := visit_repository.Instance()

	visits := repository.GetAll()

	c.JSON(http.StatusOK, visit.MarshalAll(visits))
}

type createRequest struct {
	CafeId int64 `json:"cafe_id"`
}

func Create(c *gin.Context) {
	repository := visit_repository.Instance()

	request := createRequest{}

	bindingErr := c.BindJSON(&request)

	fmt.Println(c.Request.Body)

	if bindingErr != nil {
		c.String(http.StatusBadRequest, "Invalid request")
	}

	createdVisit, createVisitErr := repository.CreateVisit(request.CafeId)

	if createVisitErr != nil {
		c.String(http.StatusBadRequest, createVisitErr.Error())
	}

	c.JSON(http.StatusOK, createdVisit.Marshal())
}
