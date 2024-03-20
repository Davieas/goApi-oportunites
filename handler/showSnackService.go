package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowSnackHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET to UaiFOOD",
	})
}
