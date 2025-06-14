package courier

import (
	"gostore/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShippingController struct {
	DB *gorm.DB
}

func NewShippingController(db *gorm.DB) *ShippingController {
	return &ShippingController{DB: db}
}

func (sc *ShippingController) AcceptTransaction(c *gin.Context) {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	courier := user.(models.User)

	var tx models.Transaction
	if err := sc.DB.First(&tx, transactionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if tx.Status != "processing" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction is not in processing state"})
		return
	}

	tx.CourierID = &courier.ID
	tx.Status = "delivering"

	if err := sc.DB.Save(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to accept transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Transaction accepted for delivery",
		"transaction_id":  tx.ID,
		"courier_assigned": courier.Name,
	})
}

func (sc *ShippingController) MarkAsArrived(c *gin.Context) {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	courier := user.(models.User)

	var tx models.Transaction
	if err := sc.DB.First(&tx, transactionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if tx.CourierID == nil || *tx.CourierID != courier.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not assigned to this transaction"})
		return
	}

	if tx.Status != "delivering" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction is not in delivering state"})
		return
	}

	tx.Status = "arrived"

	if err := sc.DB.Save(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark transaction as arrived"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Transaction marked as arrived",
		"transaction_id": tx.ID,
	})
}
