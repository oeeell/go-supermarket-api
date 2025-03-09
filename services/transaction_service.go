package services

import (
	"errors"
	"supermarket-api/config"
	"supermarket-api/models"
	"time"
)

// AddTransaction processes a transaction
func AddTransaction(transaction *models.Transaction) error {
	// Fetch product details
	var product models.Product
	if err := config.DB.First(&product, transaction.ProductID).Error; err != nil {
		return errors.New("product not found")
	}

	// Check stock availability
	if product.Stock < transaction.Quantity {
		return errors.New("not enough stock")
	}

	// Calculate total price
	transaction.Total = float64(transaction.Quantity) * product.Price

	// Deduct stock
	product.Stock -= transaction.Quantity
	config.DB.Save(&product)

	// Save transaction
	if err := config.DB.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

// GetTransactions fetches a list of transactions with pagination
func GetTransactions(limit, offset int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := config.DB.Limit(limit).Offset(offset).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetDailySalesReport returns total sales & revenue for today
func GetDailySalesReport() (int, float64, error) {
	var totalTransactions int64
	var totalRevenue float64

	startOfDay := time.Now().Truncate(24 * time.Hour)

	err := config.DB.Model(&models.Transaction{}).
		Where("created_at >= ?", startOfDay).
		Count(&totalTransactions).
		Select("COALESCE(SUM(total), 0)").
		Scan(&totalRevenue).Error

	if err != nil {
		return 0, 0, err
	}

	return int(totalTransactions), totalRevenue, nil
}

// GetMonthlySalesReport returns total sales & revenue for this month
func GetMonthlySalesReport() (int, float64, error) {
	var totalTransactions int64
	var totalRevenue float64

	startOfMonth := time.Now().Truncate(24 * time.Hour).AddDate(0, 0, -time.Now().Day()+1)

	err := config.DB.Model(&models.Transaction{}).
		Where("created_at >= ?", startOfMonth).
		Count(&totalTransactions).
		Select("COALESCE(SUM(total), 0)").
		Scan(&totalRevenue).Error

	if err != nil {
		return 0, 0, err
	}

	return int(totalTransactions), totalRevenue, nil
}

// GetBestSellingProducts returns top-selling products
func GetBestSellingProducts(limit int) ([]models.Product, error) {
	var products []models.Product

	err := config.DB.Raw(`
		SELECT p.id, p.name, SUM(t.quantity) as total_sold
		FROM products p
		JOIN transactions t ON p.id = t.product_id
		GROUP BY p.id, p.name
		ORDER BY total_sold DESC
		LIMIT ?
	`, limit).Scan(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}