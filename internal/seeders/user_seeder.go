package seeders

import (
	"log"

	"gostore/internal/models"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []models.User{
		{Name: "Admin", Email: "admin@example.com", Password: "admin123", Address: "jalan satu", Phone: "23134", Role: "admin", Balance: 0},
		{Name: "Customer One", Email: "customer1@example.com", Password: "customer123",Address: "jalan satu", Phone: "1234",  Role: "customer", Balance: 100000},
		{Name: "Courier One", Email: "courier1@example.com", Password: "courier123", Address: "jalan satu", Phone: "23134", Role: "courier", Balance: 0},
	}

	for _, user := range users {
		var existing models.User
		err := db.Where("email = ?", user.Email).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&user).Error; err != nil {
				log.Printf("Failed to seed user '%s': %v", user.Email, err)
			} else {
				log.Printf("Seeded user: %s", user.Email)
			}
		}
	}
}
