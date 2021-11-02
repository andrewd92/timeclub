package club_controller

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/club_repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func All(c *gin.Context) {
	clubRepository, err := club_repository.Instance()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "club repository instantiate err")
	}

	clubs := clubRepository.GetAll()

	c.JSON(http.StatusOK, club.MarshalAll(clubs))
}
