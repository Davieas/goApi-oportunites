package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSnackHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE to UaiFOOD",
	})
}