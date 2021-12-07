package visit_controller

import (
	"github.com/andrewd92/timeclub/visit_service/application/visit_service"
	"github.com/andrewd92/timeclub/visit_service/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	timePkg "time"
)

const clubIdKey = "club_id"

func All(c *gin.Context) {
	now := timePkg.Now()
	getForTime(now, c)
}

func ForTime(c *gin.Context) {
	loc, _ := timePkg.LoadLocation("Europe/Warsaw")
	timeStr := c.Param("time")
	time, err := timePkg.ParseInLocation(utils.TimeFormat, timeStr, loc)
	if err != nil {
		log.WithError(err).WithField("time", timeStr).Error("can not parse requested time for visits")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	getForTime(time, c)
}

func getForTime(time timePkg.Time, c *gin.Context) {
	clubId, err := strconv.ParseInt(c.Param(clubIdKey), 10, 64)

	if err != nil {
		log.WithError(err).WithField(clubIdKey, c.Param(clubIdKey)).Error("can not parse club id")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	service := visit_service.Instance()

	response, responseErr := service.All(clubId, time)

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

	now := timePkg.Now()

	response, createVisitErr := service.Create(request.ClubId, request.CardId, now)

	if createVisitErr != nil {
		c.String(http.StatusInternalServerError, createVisitErr.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
