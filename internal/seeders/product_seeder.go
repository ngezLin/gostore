package seeders

import (
	"log"

	"gostore/internal/models"

	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) {
	products := []models.Product{
		{Name: "Smartphone", Price: 5000000, Stock: 50, ProductCategoryID: 1},
		{Name: "T-Shirt", Price: 100000, Stock: 100, ProductCategoryID: 2},
		{Name: "Novel", Price: 75000, Stock: 60, ProductCategoryID: 3},
		{Name: "Instant Noodles", Price: 3000, Stock: 200, ProductCategoryID: 4},
		{Name: "Office Chair", Price: 800000, Stock: 20, ProductCategoryID: 5},
	}

	for _, product := range products {
		var existing models.Product
		err := db.Where("name = ?", product.Name).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&product).Error; err != nil {
				log.Printf("Failed to seed product '%s': %v", product.Name, err)
			} else {
				log.Printf("Seeded product: %s", product.Name)
			}
		}
	}
}
