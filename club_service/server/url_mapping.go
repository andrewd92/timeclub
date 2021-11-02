package server

import (
	"github.com/andrewd92/timeclub/club_service/controller/club_controller"
)

func mapUrls() {
	router.GET("/clubs", club_controller.All)
}
