package controllers

import (
	"net/http"
	"strconv"

	"sanbercode-golang-final-project-nando/pkg/database"
	"sanbercode-golang-final-project-nando/pkg/repository"
	"sanbercode-golang-final-project-nando/pkg/structs"

	"github.com/gin-gonic/gin"
)

func GetAllCarts(c *gin.Context) {
	username, _, _ := c.Request.BasicAuth()
	carts, err := repository.GetAllCarts(database.DBConnection, username)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": carts,
	})
}

func AddToCart(c *gin.Context) {
	var newCart structs.Cart
	var result string

	err := c.ShouldBindJSON(&newCart)
	if err != nil {
		panic(err)
	}

	if _, errMenu := repository.GetMenu(database.DBConnection, newCart.MenuId); errMenu != nil {
		result = "Menu not found!"
	} else {
		username, _, _ := c.Request.BasicAuth()
		newCart.OrderId = username

		if newCart.CartQuantity <= 0 {
			result = "Quantity must greater than 0!"
		} else {
			if oldCart, errBook := repository.GetCart(database.DBConnection, newCart); errBook != nil {
				err = repository.AddToCart(database.DBConnection, newCart)
				result = "New menu was successfully added!"
			} else {
				oldCart.CartQuantity += newCart.CartQuantity
				err = repository.UpdateCart(database.DBConnection, oldCart)
				result = "Menu was successfully updated!"
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func DeleteFromCart(c *gin.Context) {
	var cart structs.Cart
	var result string

	err := c.ShouldBindJSON(&cart)
	if err != nil {
		panic(err)
	}

	if _, errMenu := repository.GetMenu(database.DBConnection, cart.MenuId); errMenu != nil {
		result = "Menu not found!"
	} else {
		username, _, _ := c.Request.BasicAuth()
		cart.OrderId = username

		if oldCart, errBook := repository.GetCart(database.DBConnection, cart); errBook != nil {
			result = "Menu not found in your cart!"
		} else {
			updateQuantity := oldCart.CartQuantity - cart.CartQuantity

			if updateQuantity < 0 {
				result = "Deleted menus exceed those in the cart!"
			} else {
				if updateQuantity == 0 {
					err = repository.DeleteFromCart(database.DBConnection, oldCart)
					result = "Menu was successfully deleted!"
				} else {
					oldCart.CartQuantity -= cart.CartQuantity
					err = repository.UpdateCart(database.DBConnection, oldCart)
					result = "Menu was successfully updated!"
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func ClearCart(c *gin.Context) {
	username, _, _ := c.Request.BasicAuth()
	var result string

	err := repository.ClearCart(database.DBConnection, username)
	if err != nil {
		result = err.Error()
	} else {
		result = "All orders in cart successfully deleted!"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func CheckoutCart(c *gin.Context) {
	username, _, _ := c.Request.BasicAuth()
	carts, _ := repository.GetAllCarts(database.DBConnection, username) // GET ALL CARTS

	var listDtrans []structs.Dtrans = []structs.Dtrans{}
	var result string

	if len(carts) == 0 {
		result = "Failed to checkout order, your cart is empty!"
	} else {
		htrans, _ := repository.GetOrder(database.DBConnection, username) // GET HTRANS
		subtotal := 0

		// LOOPING TO INSERT PER ORDER TO DTRANS
		for _, cart := range carts {
			menu, _ := repository.GetMenu(database.DBConnection, cart.MenuId)

			dtrans_total := cart.CartQuantity * menu.MenuPrice
			subtotal += dtrans_total

			var dtrans structs.Dtrans = structs.Dtrans{
				OrderId:        cart.OrderId,
				MenuId:         cart.MenuId,
				DtransQuantity: cart.CartQuantity,
				DtransTotal:    strconv.Itoa(dtrans_total),
			}

			err := repository.InsertDtrans(database.DBConnection, dtrans)

			if err != nil {
				panic(err)
			}

			listDtrans = append(listDtrans, dtrans)
		}

		// AFTER INSERT TO DTRANS, CLEAR ALL ORDER IN CART
		repository.ClearCart(database.DBConnection, username)

		// UPDATE HTRANS
		htrans.HtransSubtotal = strconv.Itoa(subtotal)
		err := repository.UpdateOrder(database.DBConnection, htrans)

		if err != nil {
			result = "Failed to checkout order!"
		} else {
			result = "Order successfully checkout!"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
		"data":    listDtrans,
	})
}
