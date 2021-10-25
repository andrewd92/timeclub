package club_controller

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/club_repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func All(c *gin.Context) {
	clubRepository := club_repository.NewClubInMemoryRepository()

	clubs := clubRepository.GetAll()

	c.JSON(http.StatusOK, club.MarshalAll(clubs))
}
