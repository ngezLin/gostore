package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if userRole != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient role"})
			return
		}

		c.Next()
	}
}
