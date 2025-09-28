package handlers

import (
	"car_inventory/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
	// Fetch cars from database
	c.JSON(http.StatusOK, gin.H{"message": "Get all cars"})
}

func AddCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add car to database
	c.JSON(http.StatusOK, gin.H{"message": "Car added", "car": car})
}
