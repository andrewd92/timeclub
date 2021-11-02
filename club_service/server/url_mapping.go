package server

import (
	"github.com/andrewd92/timeclub/club_service/controller/card_controller"
	"github.com/andrewd92/timeclub/club_service/controller/club_controller"
)

func mapUrls() {
	router.GET("/clubs", club_controller.All)
	router.POST("/card", card_controller.Create)
	router.GET("/card/templates", card_controller.Templates)
	router.POST("/card/template", card_controller.CreateTemplate)
}
