package main

import (
	"github.com/gin-gonic/gin"
	"github.com/testmock/controllers"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.PingController)
	router.Run(":8080")
}
