package controllers

import (
	"net/http"
	"strconv"
	"supermarket-api/models"
	"supermarket-api/services"

	"github.com/gin-gonic/gin"
)

// Get all products with pagination
func GetProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	products, err := services.GetPaginatedProducts(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Create product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.AddProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}
