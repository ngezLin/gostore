package customer

import (
	"net/http"

	"gostore/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DetailController struct {
	DB *gorm.DB
}

func NewDetailController(db *gorm.DB) *DetailController {
	return &DetailController{DB: db}
}

// Get customer detail by current user
func (dc *DetailController) GetCustomerDetail(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser := user.(models.User)

	var detail models.CustomerDetail
	if err := dc.DB.Preload("User").Where("user_id = ?", currentUser.ID).First(&detail).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer detail not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": detail})
}
// UpdateBalance updates the user's balance in their customer detail.
func (dc *DetailController) UpdateBalance(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser := user.(models.User)

	var input struct {
		UserBalance float64 `json:"user_balance"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var detail models.CustomerDetail
	if err := dc.DB.Preload("User").Where("user_id = ?", currentUser.ID).First(&detail).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer detail not found"})
		return
	}

	detail.UserBalance = input.UserBalance
	if err := dc.DB.Save(&detail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Balance updated", "data": detail})
}
// 