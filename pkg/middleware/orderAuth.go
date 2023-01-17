package middleware

import (
	"net/http"
	"sanbercode-golang-final-project-nando/pkg/database"
	"sanbercode-golang-final-project-nando/pkg/repository"

	"github.com/gin-gonic/gin"
)

// Middleware untuk check apakah customer sudah melakukan checkout atau belum
func AuthOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _, _ := c.Request.BasicAuth()
		order, _ := repository.GetOrder(database.DBConnection, username)

		if order.HtransStatus == 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "This order has been completed!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
