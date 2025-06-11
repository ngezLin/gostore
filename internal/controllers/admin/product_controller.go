package admin

import (
	"gostore/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	if err := pc.DB.Preload("Category").First(&product, product.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "data": product})
}


func (pc *ProductController) GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := pc.DB.Preload("Category").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Find the existing product
	if err := pc.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Bind the updated data
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated product
	if err := pc.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Preload the category again
	if err := pc.DB.Preload("Category").First(&product, product.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated", "data": product})
}


func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := pc.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
