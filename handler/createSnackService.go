package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSnackHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST to UaiFOOD",
	})
}
