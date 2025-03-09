package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}
