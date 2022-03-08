package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
