package routes

import (
	"supermarket-api/controllers"
	"supermarket-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
    {
        api.GET("/products", controllers.GetProducts)
        api.POST("/products", middleware.ValidateProductInput(), controllers.CreateProduct)

        api.POST("/transactions", controllers.CreateTransaction)
        api.GET("/transactions", controllers.GetTransactions)

        api.GET("/reports/daily", controllers.GetDailySalesReport)
        api.GET("/reports/monthly", controllers.GetMonthlySalesReport)
        api.GET("/reports/best-selling", controllers.GetBestSellingProducts)
    }
}

