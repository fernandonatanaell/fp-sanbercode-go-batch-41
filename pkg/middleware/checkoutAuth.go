package middleware

import (
	"net/http"
	"sanbercode-golang-final-project-nando/pkg/database"
	"sanbercode-golang-final-project-nando/pkg/repository"

	"github.com/gin-gonic/gin"
)

// Middleware untuk check apakah customer sudah melakukan checkout atau belum
func AuthCheckout() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _, ok := c.Request.BasicAuth()
		order, err := repository.GetOrder(database.DBConnection, username)

		if !ok || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		if order.OrderId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Order not found!",
			})
			c.Abort()
			return
		} else if order.HtransStatus == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "This order isn't finished yet!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
