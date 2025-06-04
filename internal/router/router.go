package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/", func(c *gin.Context){
		if db != nil {
			c.JSON(200, gin.H{
				"message": "Database connection is successful",
			})
		} else {
			c.JSON(500, gin.H{
				"message": "Database connection failed",
			})
		}
	})
}