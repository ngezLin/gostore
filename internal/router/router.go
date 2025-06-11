package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gostore/internal/controllers"
	"gostore/internal/controllers/admin"
	"gostore/internal/middlewares"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Base API Group
	api := r.Group("/api")

	// AUTH ROUTES (Public)
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.CustomerRegister(db))
		auth.POST("/login", controllers.Login(db))
	}

	// PROTECTED ROUTES (All authenticated users)
	protected := api.Group("/")
	protected.Use(middlewares.AuthMiddleware(db))
	{
		// /api/me
		protected.GET("/me", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(200, gin.H{"user": user})
		})
	}

	// ADMIN ROUTES (Under /api/admin)
	adminGroup := api.Group("/admin")
	adminGroup.Use(middlewares.AuthMiddleware(db), middlewares.RoleMiddleware("admin"))
	{
		productController := admin.NewProductController(db)
		categoryController := admin.NewCategoryController(db)

		adminGroup.POST("/products", productController.CreateProduct)
		adminGroup.GET("/products", productController.GetAllProducts)
		adminGroup.PUT("/products/:id", productController.UpdateProduct)
		adminGroup.DELETE("/products/:id", productController.DeleteProduct)


		adminGroup.POST("/categories", categoryController.CreateCategory)
		adminGroup.GET("/categories", categoryController.GetAllCategories)
		adminGroup.PUT("/categories/:id", categoryController.UpdateCategory)
		adminGroup.DELETE("/categories/:id", categoryController.DeleteCategory)
	}
}
