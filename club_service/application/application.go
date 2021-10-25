package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Err: ", err.Error())
	}
}
