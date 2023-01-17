package middleware

import (
	"net/http"
	"sanbercode-golang-final-project-nando/pkg/database"
	"sanbercode-golang-final-project-nando/pkg/repository"

	"github.com/gin-gonic/gin"
)

// Middleware untuk check apakah benar customer?
func AuthCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _, ok := c.Request.BasicAuth()
		_, err := repository.GetOrder(database.DBConnection, username)

		if !ok || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
