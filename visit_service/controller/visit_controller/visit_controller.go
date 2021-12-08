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

type VisitController struct {
	visitService visit_service.VisitService
}

func (vc VisitController) All(c *gin.Context) {
	now := timePkg.Now()
	vc.getForTime(now, c)
}

func (vc VisitController) ForTime(c *gin.Context) {
	timeStr := c.Param("time")
	time, err := vc.parseTime(timeStr)
	if err != nil {
		log.WithError(err).WithField("time", timeStr).Error("can not parse requested time for visits")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	vc.getForTime(time, c)
}

func (vc VisitController) parseTime(timeStr string) (timePkg.Time, error) {
	loc, _ := timePkg.LoadLocation("Europe/Warsaw")
	return timePkg.ParseInLocation(utils.TimeFormat, timeStr, loc)
}

func (vc VisitController) getForTime(time timePkg.Time, c *gin.Context) {
	clubId, err := strconv.ParseInt(c.Param(clubIdKey), 10, 64)

	if err != nil {
		log.WithError(err).WithField(clubIdKey, c.Param(clubIdKey)).Error("can not parse club id")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, responseErr := vc.visitService.All(clubId, time)

	if nil != responseErr {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

type createRequest struct {
	ClubId int64 `json:"club_id"`
	CardId int64 `json:"card_id"`
}

func (vc VisitController) Create(c *gin.Context) {
	request := createRequest{}

	bindingErr := c.BindJSON(&request)

	if bindingErr != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	now := timePkg.Now()

	response, createVisitErr := vc.visitService.Create(request.ClubId, request.CardId, now)

	if createVisitErr != nil {
		c.String(http.StatusInternalServerError, createVisitErr.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
