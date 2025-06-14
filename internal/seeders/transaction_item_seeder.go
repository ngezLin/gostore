package seeders

import (
	"log"

	"gostore/internal/models"

	"gorm.io/gorm"
)

func SeedTransactionItems(db *gorm.DB) {
	items := []models.TransactionItem{
		{TransactionID: 1, ProductID: 1, Quantity: 1, SubTotal: 5000000},
		{TransactionID: 1, ProductID: 2, Quantity: 2, SubTotal: 200000},
	}

	for _, item := range items {
		if err := db.Create(&item).Error; err != nil {
			log.Printf("Failed to seed transaction item: %v", err)
		} else {
			log.Printf("Seeded transaction item for Transaction ID: %d", item.TransactionID)
		}
	}
}
