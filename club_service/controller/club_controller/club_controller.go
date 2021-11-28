package club_controller

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/club_repository"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func All(c *gin.Context) {
	clubs, getErr := club_repository.Instance().GetAll()

	if getErr != nil {
		log.WithError(getErr).Error("Can not find clubs")
		c.String(http.StatusInternalServerError, "DB error")
		return
	}

	c.JSON(http.StatusOK, club.MarshalAll(clubs))
}
