package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"weather-station-api/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/", controllers.GetAllStations)
	router.GET("/:id", controllers.GetStationById)

	router.Use(cors.Default())
	router.Run(":3000")
}