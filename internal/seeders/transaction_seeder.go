package seeders

import (
	"log"

	"gostore/internal/models"

	"gorm.io/gorm"
)

func SeedTransactions(db *gorm.DB) {
	tx := models.Transaction{
		CustomerID: 2, // assuming "Customer One"
		CourierID:  nil,
		TotalAmount: 150000,
		Status:      "processing",
	}

	if err := db.Create(&tx).Error; err != nil {
		log.Printf("Failed to seed transaction: %v", err)
	} else {
		log.Printf("Seeded transaction with ID: %d", tx.ID)
	}
}
