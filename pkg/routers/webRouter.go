package routers

import (
	"sanbercode-golang-final-project-nando/pkg/controllers"
	"sanbercode-golang-final-project-nando/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// === ADMIN ===
	admin := router.Group("/admin").Use(middleware.AuthAdmin())

	// admin - order (htrans)
	admin.GET("/order", controllers.GetOrderFromAllUser)        // SHOW ORDER FROM ALL CUSTOMER
	admin.POST("/order", controllers.CreateNewOrder)            // CREATE NEW ORDER
	admin.GET("/order/detail", controllers.GetOrderFromGivenID) // SHOW DETAIL ORDER FROM GIVEN ID

	// admin - menu
	admin.GET("/menu", controllers.GetAllMenuWithTrashed) // SHOW ALL MENU WITH TRASHED
	admin.POST("/menu", controllers.CreateNewMenu)        // CREATE NEW MENU
	admin.PUT("/menu", controllers.UpdateMenu)            // UPDATE MENU
	admin.DELETE("/menu", controllers.DeleteMenu)         // DELETE MENU (SOFT DELETE)

	// === CUSTOMER === (BLM KERJA)
	customer := router.Group("/customer").Use(middleware.AuthCustomer())

	// customer - order (htrans)
	order := customer.Use(middleware.AuthOrder())
	order.GET("/order", controllers.GetAllCarts)           // SHOW CUSTOMER CART
	order.POST("/order/insert", controllers.AddToCart)     // INSERT TO CART
	order.PUT("/order/delete", controllers.DeleteFromCart) // UPDATE OR DELETE FROM CART
	order.POST("/order/clear", controllers.ClearCart)      // CLEAR CART
	order.POST("/checkout", controllers.CheckoutCart)      // SUBMIT CHECKOUT ORDER

	// customer - checkout
	checkout := router.Group("/checkout").Use(middleware.AuthCheckout())
	checkout.GET("/", controllers.GetDtransFromCurrentUser) // SHOW CUSTOMER CHECKOUT ORDERS

	// === ALL USER CAN ACCESS THIS ===
	router.GET("/menu", controllers.GetAllMenu)        // SHOW ALL MENU WITHOUT TRASHED
	router.GET("/menu/search", controllers.SearchMenu) // SEARCH MENU BY NAME

	return router
}
