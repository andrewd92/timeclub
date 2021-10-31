package visit_controller

import (
	"github.com/andrewd92/timeclub/visit_service/application/visit_service"
	"github.com/andrewd92/timeclub/visit_service/domain/price_list"
	"github.com/andrewd92/timeclub/visit_service/domain/price_list/price"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/repository/visit_repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func All(c *gin.Context) {
	repository := visit_repository.Instance()

	visits := repository.GetAll()

	response, responseErr := visit.MarshalAll(visits, price_list.DefaultPriceList(), price.USD())

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
	service := visit_service.Instance()

	request := createRequest{}

	bindingErr := c.BindJSON(&request)

	if bindingErr != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	response, createVisitErr := service.Create(request.ClubId, request.CardId)

	if createVisitErr != nil {
		c.String(http.StatusInternalServerError, createVisitErr.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
