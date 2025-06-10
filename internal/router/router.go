package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gostore/internal/controllers"
	"gostore/internal/middlewares"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	// AUTH
	auth := api.Group("/auth")
	auth.POST("/register", controllers.CustomerRegister(db))
	auth.POST("/login", controllers.Login(db))

	// PROTECTED ROUTES
	protected := api.Group("/")
	protected.Use(middlewares.AuthMiddleware(db))

	// Test route to get current user
	protected.GET("/me", func(c *gin.Context) {
		user, _ := c.Get("user")
		c.JSON(200, gin.H{"user": user})
	})

	// ADMIN ONLY (optional)
	// admin := protected.Group("/admin")
	// admin.Use(middlewares.RoleMiddleware("admin"))
	// admin.GET("/dashboard", ...)

	// CUSTOMER ONLY (optional)
	// customer := protected.Group("/customer")
	// customer.Use(middlewares.RoleMiddleware("customer"))
	// customer.GET("/transactions", ...)

	// COURIER ONLY (optional)
	// courier := protected.Group("/courier")
	// courier.Use(middlewares.RoleMiddleware("courier"))
	// courier.GET("/deliveries", ...)
}
