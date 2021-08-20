package server

import (
	"github.com/gin-gonic/gin"
	"t-rex_wrapper/handler"
)

func GoGinServer() {
	server := gin.Default()
	server.GET("/testing", handler.StartPage)
	// For each matched request Context will hold the route definition
	server.POST("/fan", handler.SetFanPercent)
	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "c")
}
