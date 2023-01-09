package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"weather-station-api/services"
)

// This is the Controller for the Stations architect logic.
// It sends the data to the app matching the business logic.

func GetAllStations(c *gin.Context) {
	stations, err := services.GetStations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stations)
}

func GetStationById(c *gin.Context) {
	id := c.Params.ByName("id")
	station, err := services.GetStationDetails(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, station)
}