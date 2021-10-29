package server

import (
	"github.com/andrewd92/timeclub/club_service/controller/card_controller"
	"github.com/andrewd92/timeclub/club_service/controller/club_controller"
	"github.com/andrewd92/timeclub/club_service/controller/visit_controller"
)

func mapUrls() {
	router.GET("/clubs", club_controller.All)
	router.GET("/visits", visit_controller.All)
	router.POST("/visit", visit_controller.Create)
	router.POST("/card", card_controller.Create)
	router.GET("/card/min/:clubId", card_controller.MinAvailable)
	router.GET("/card/templates", card_controller.Templates)
}
