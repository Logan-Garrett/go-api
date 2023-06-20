package main

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/Jobs", controllers.GetJobs)
	router.GET("/Jobs/:id", controllers.GetJob)
	router.PATCH("/Jobs/:id", controllers.ModifyJob)
	router.POST("/Jobs", controllers.AddJob)
	router.Run("localhost:9090")
}
