package main

import (
	"car_inventory/config"
	"car_inventory/handlers"
	"car_inventory/middlewares"
	"car_inventory/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	err = db.AutoMigrate(&models.Car{})
	if err != nil {
		panic("Failed to migrate database")
	}
	log.Println("Database connected and migrated")
	r := gin.Default()

	// Apply middlewares
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.AuthMiddleware())

	// Routes
	r.GET("/cars", handlers.GetCars)
	r.POST("/cars", handlers.AddCar)

	// Start server
	r.Run(":8080")
}
