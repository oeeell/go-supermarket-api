package services

import (
	"supermarket-api/models"
	"supermarket-api/repositories"
)

// Get products with pagination
func GetPaginatedProducts(limit, offset int) ([]models.Product, error) {
	return repositories.GetAllProducts(limit, offset)
}

// Create new product
func AddProduct(product *models.Product) error {
	return repositories.CreateProduct(product)
}
