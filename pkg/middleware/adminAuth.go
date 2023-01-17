package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware untuk check apakah benar admin?
func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || username != "admin" || password != "nimda" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
