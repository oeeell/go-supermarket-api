package controllers

import (
	"net/http"
	"supermarket-api/models"
	"supermarket-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTransaction handles checkout transactions
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call the service
	err := services.AddTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// GetTransactions lists all transactions
func GetTransactions(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	transactions, err := services.GetTransactions(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// GetDailySalesReport API
func GetDailySalesReport(c *gin.Context) {
	totalTransactions, totalRevenue, err := services.GetDailySalesReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_transactions": totalTransactions,
		"total_revenue":      totalRevenue,
	})
}

// GetMonthlySalesReport API
func GetMonthlySalesReport(c *gin.Context) {
	totalTransactions, totalRevenue, err := services.GetMonthlySalesReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_transactions": totalTransactions,
		"total_revenue":      totalRevenue,
	})
}

// GetBestSellingProducts API
func GetBestSellingProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	products, err := services.GetBestSellingProducts(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
