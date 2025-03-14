package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name" gorm:"uniqueIndex"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
