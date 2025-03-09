package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateProductInput checks if the required fields are present in the request
func ValidateProductInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Name  string  `json:"name"`
			Price float64 `json:"price"`
			Stock int     `json:"stock"`
		}

		// Print raw request body for debugging
		body, _ := c.GetRawData()
		fmt.Println("Raw Request Body:", string(body))

		// Rebind the request after reading raw data
		c.Request.Body = http.NoBody // Reset body for parsing
		if err := c.ShouldBindJSON(&input); err != nil {
			fmt.Println("Bind error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
			c.Abort()
			return
		}

		fmt.Printf("Received JSON: %+v\n", input) // Debug

		if input.Name == "" || input.Price <= 0 || input.Stock < 0 {
			fmt.Println("Validation failed:", input)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
			c.Abort()
			return
		}

		c.Next()
	}
}
