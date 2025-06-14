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

type BalanceRequest struct {
	Amount float64 `json:"amount"`
}

func (dc *DetailController) UpdateBalance(c *gin.Context) {
	var req BalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	customer := user.(models.User)
	customer.Balance += req.Amount

	if err := dc.DB.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Balance updated successfully",
		"balance": customer.Balance,
	})
}
