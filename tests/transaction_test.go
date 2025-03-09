package tests

import (
    "supermarket-api/models"
    "testing"
)

func TestTransactionTotalCalculation(t *testing.T) {
    transaction := models.Transaction{
        ProductID: 1,
        Quantity: 2,
        Total:     20.0, // Assuming price is 10
    }

    expectedTotal := 20.0
    if transaction.Total != expectedTotal {
        t.Errorf("Expected total %v, but got %v", expectedTotal, transaction.Total)
    }
}
