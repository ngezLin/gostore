package seeders

import (
	"gostore/internal/models"
	"log"

	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {
	categories := []models.ProductCategory{
		{Name: "Electronics"},
		{Name: "Clothing"},
		{Name: "Books"},
		{Name: "Food"},
		{Name: "Furniture"},
	}

	for _, category := range categories {
		var existing models.ProductCategory
		err := db.Where("name = ?", category.Name).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Failed to seed category '%s': %v", category.Name, err)
			} else {
				log.Printf("Seeded category: %s", category.Name)
			}
		}
	}
}