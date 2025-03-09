package main

import (
	"log"
	"supermarket-api/config"
	"supermarket-api/middleware"
	"supermarket-api/models"
	"supermarket-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{}, &models.Transaction{})

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())

	routes.SetupRoutes(router)

	log.Println("🚀 Server running on port 8080")
	router.Run(":8080")
}
