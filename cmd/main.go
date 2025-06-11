package main

import (
	"gostore/config"
	"gostore/internal/models"
	"gostore/internal/router"
	"gostore/internal/seeders"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	db := config.ConnectDatabase()

	db.AutoMigrate(
		&models.User{},
		&models.CustomerDetail{},
		&models.Product{},
		models.ProductCategory{},
		&models.Transaction{},
		&models.TransactionItem{},
	)
	//seeders
	seeders.SeedCategories(db)

	router.SetupRoutes(r, db)
	r.Run(":8080")
}