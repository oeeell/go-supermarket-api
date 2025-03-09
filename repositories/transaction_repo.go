package repositories

import (
	"supermarket-api/config"
	"supermarket-api/models"
)

// GetProductByID fetches a product by ID
func GetProductByID(productID uint, product *models.Product) error {
	return config.DB.First(product, productID).Error
}

// UpdateProduct updates a product in the database
func UpdateProduct(product *models.Product) error {
	return config.DB.Save(product).Error
}

// CreateTransaction saves a new transaction
func CreateTransaction(transaction *models.Transaction) error {
	return config.DB.Create(transaction).Error
}

// GetAllTransactions retrieves all transactions with pagination
func GetAllTransactions(limit, offset int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := config.DB.Limit(limit).Offset(offset).Find(&transactions).Error
	return transactions, err
}
