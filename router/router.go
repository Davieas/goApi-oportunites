package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	// Initialize Router
	r := gin.Default()
	//Initializing Routes
	initializeRoutes(r)
	//Running Server
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
