package server

import (
	"github.com/andrewd92/timeclub/card_service/controller/card_controller"
)

func mapUrls() {
	router.GET("/card", card_controller.All)
	router.POST("/card", card_controller.Create)
	router.GET("/card/templates", card_controller.Templates)
	router.POST("/card/template", card_controller.CreateTemplate)
}
