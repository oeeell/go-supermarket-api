package repositories

import (
	"supermarket-api/config"
	"supermarket-api/models"
)

// Fetch all products with pagination
func GetAllProducts(limit, offset int) ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

// Create new product
func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}
