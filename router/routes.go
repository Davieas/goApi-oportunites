package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//Show Services
		v1.GET("/snacksService", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Welcome to UaiFOOD",
			})
		})
		v1.POST("/snacksService", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "HELLO to UaiFOOD",
			})
		})
		v1.DELETE("/snacksService", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "BYE to UaiFOOD",
			})
		})
		v1.PUT("/snacksService", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "NEW to UaiFOOD",
			})
		})
		v1.GET("/snacksServices", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Welcome to all new snacks Services",
			})
		})
	}
}
