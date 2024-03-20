package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateSnackHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "UPDATE to UaiFOOD",
	})
}
