package main

import (
	"gostore/config"
	"gostore/internal/router"
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
		
	)

	router.SetupRoutes(r, db)
	r.Run(":8080")
}