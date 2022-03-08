package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//RegisterAPIRoutes 注册API路由
func RegisterAPIRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "v1"})
		})
	}

	v2 := router.Group("v2")
	{
		v2.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "v2"})
		})
	}
}
