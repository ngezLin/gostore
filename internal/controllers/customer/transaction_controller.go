package customer

import (
	"net/http"

	"gostore/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionController struct {
	DB *gorm.DB
}

func NewTransactionController(db *gorm.DB) *TransactionController {
	return &TransactionController{DB: db}
}

type TransactionRequest struct {
	Items []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	customer := user.(models.User)

	var totalAmount float64
	var items []models.TransactionItem

	// for _, i := range req.Items {
	// 	var product models.Product
	// 	if err := tc.DB.Preload("Category").First(&product, i.ProductID).Error; err != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	// 		return
	// 	}

	// 	subTotal := product.Price * float64(i.Quantity)
	// 	totalAmount += subTotal

	// 	items = append(items, models.TransactionItem{
	// 		ProductID: i.ProductID,
	// 		Quantity:  i.Quantity,
	// 		SubTotal:  subTotal,
	// 	})
	// }

	for _, i := range req.Items {
		var product models.Product
		if err := tc.DB.Preload("Category").First(&product, i.ProductID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		//Check stock availability
		if product.Stock < i.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Insufficient stock for product: " + product.Name,
			})
			return
		}

		//Decrease stock
		product.Stock -= i.Quantity
		if err := tc.DB.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
			return
		}

		subTotal := product.Price * float64(i.Quantity)
		totalAmount += subTotal

		items = append(items, models.TransactionItem{
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			SubTotal:  subTotal,
		})
	}


	if customer.Balance < totalAmount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	customer.Balance -= totalAmount
	if err := tc.DB.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update balance"})
		return
	}


	tx := models.Transaction{
		CustomerID:  customer.ID,
		TotalAmount: totalAmount,
		Status:      "processing",
		Items:       items,
	}

	if err := tc.DB.Create(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Reload transaction with full data
	if err := tc.DB.Preload("Customer").
		Preload("Courier").
		Preload("Items.Product").
		First(&tx, tx.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load transaction data"})
		return
	}

	// Format response items
	var responseItems []gin.H
	for _, item := range tx.Items {
		responseItems = append(responseItems, gin.H{
			"name":     item.Product.Name,
			"quantity": item.Quantity,
		})
	}

	// Format courier info if available
	var courierInfo any
	if tx.Courier != nil {
		courierInfo = gin.H{
			"name":  tx.Courier.Name,
			"phone": tx.Courier.Phone,
		}
	}

	// Return simplified custom response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Transaction created",
		"transaction": gin.H{
			"items": responseItems,
			"customer": gin.H{
				"name":    tx.Customer.Name,
				"phone":   tx.Customer.Phone,
				"address": tx.Customer.Address,
			},
			"courier":  courierInfo,
			"subtotal": tx.TotalAmount,
		},
	})
}
