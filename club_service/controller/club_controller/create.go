package club_controller

import (
	"bytes"
	"github.com/andrewd92/timeclub/club_service/api/http/create"
	"github.com/andrewd92/timeclub/club_service/application/club_service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Create(c *gin.Context) {
	request := create.Request{}

	bindingErr := c.BindJSON(&request)

	if bindingErr != nil {
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(c.Request.Body)
		if err != nil {
			log.WithError(err).Error("can not read request body")
		}
		requestBody := buf.String()

		log.WithError(bindingErr).WithField("request", requestBody).Error("can not bind request body")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	service := club_service.Instance()

	response, err := service.Create(request)

	if err != nil {
		log.WithError(err).WithField("request", request).Error("can not create club")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.WithField("response", response).Debug("club created")

	c.JSON(http.StatusOK, response)
}
