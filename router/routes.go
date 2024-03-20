package router

import (

	"github.com/Davieas/goapi-oportunites/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//Show Services
		v1.GET("/snacksService", handler.ShowSnackHandler)
		v1.POST("/snacksService", handler.CreateSnackHandler)
		v1.DELETE("/snacksService", handler.DeleteSnackHandler)
		v1.PUT("/snacksService", handler.UpdateSnackHandler)
		v1.GET("/snacksServices", handler.ListSnackHandler)
		
	}
}
