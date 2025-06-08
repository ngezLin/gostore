package router

import (
	"gostore/internal/controllers"
	"gostore/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())

	api.GET("/me", authController.Profile)
	
	// Admin-only routes
	adminRoutes := api.Group("/admin")
	adminRoutes.Use(middlewares.RoleMiddleware("admin"))
	{
		// TODO: add product/category CRUD
	}

	// Customer-only routes
	customerRoutes := api.Group("/customer")
	customerRoutes.Use(middlewares.RoleMiddleware("customer"))
	{
		// TODO: add cart, checkout
	}

	// Courier-only routes
	courierRoutes := api.Group("/courier")
	courierRoutes.Use(middlewares.RoleMiddleware("courier"))
	{
		// TODO: add transaction delivery status update
	}
}
