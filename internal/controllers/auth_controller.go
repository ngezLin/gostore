package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"

	"gostore/internal/models"
	"gostore/pkg/utils"
)

// Register handler — supports all roles during development
func CustomerRegister(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if user already exists
		var existing models.User
		if err := db.Where("email = ?", input.Email).First(&existing).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
			return
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		user := models.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: string(hashedPassword),
			Address:  input.Address,
			Phone:    input.Phone,
			Role:     input.Role,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// If customer, create associated detail
		if user.Role == "customer" {
			customerDetail := models.CustomerDetail{
				UserID:      user.ID,
				UserBalance: 0,
			}
			if err := db.Create(&customerDetail).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer detail"})
				return
			}
		}

		c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
	}
}

// Login handler
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Find user
		var user models.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Compare password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Generate token
		token, err := utils.GenerateJWT(user.ID, user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": `Login successful`,
			"token": token,
			"user": gin.H{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
				"role":  user.Role,
			},
		})
	}
}
