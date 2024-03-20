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
func ShowSnackHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET to UaiFOOD",
	})
}
func DeleteSnackHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE to UaiFOOD",
	})
}
func UpdateSnackHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "UPDATE to UaiFOOD",
	})
}
func ListSnackHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GETALL to UaiFOOD",
	})
}
