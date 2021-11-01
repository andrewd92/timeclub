package server

import (
	"fmt"
	"github.com/andrewd92/timeclub/visit_service/config"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	err := router.Run(":" + config.Instance().Server.Port.Http)
	if err != nil {
		fmt.Println("Err: ", err.Error())
	}
}
