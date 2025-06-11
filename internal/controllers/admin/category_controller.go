package admin

import (
	"gostore/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{DB: db}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var category models.ProductCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "data": category})
}

func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	var categories []models.ProductCategory
	if err := cc.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var category models.ProductCategory
	id := c.Param("id")

	if err := cc.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated", "data": category})
}

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := cc.DB.Delete(&models.ProductCategory{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
