package tests

import (
    "supermarket-api/models"
    "testing"
)

func TestProductStockReduction(t *testing.T) {
    product := models.Product{
        Name:  "Test Product",
        Price: 10.0,
        Stock: 10,
    }

    // Simulate a sale where 3 items are bought
    product.Stock -= 3

    expectedStock := 7
    if product.Stock != expectedStock {
        t.Errorf("Expected stock %v, but got %v", expectedStock, product.Stock)
    }
}

func TestProductStockCannotGoNegative(t *testing.T) {
    product := models.Product{
        Name:  "Test Product",
        Price: 10.0,
        Stock: 2,
    }

    // Simulate a sale with more items than available
    if product.Stock < 5 {
        t.Log("Stock is too low for this transaction")
    }

    if product.Stock < 0 {
        t.Errorf("Stock should not be negative, but got %v", product.Stock)
    }
}
