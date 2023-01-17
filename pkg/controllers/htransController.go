package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"sanbercode-golang-final-project-nando/pkg/database"
	"sanbercode-golang-final-project-nando/pkg/repository"

	"github.com/gin-gonic/gin"
)

// ORDER (HTRANS)
func GetOrderFromGivenID(c *gin.Context) {
	var result gin.H
	order_id := c.Query(("order_id"))

	if orders, errOrder := repository.GetOrder(database.DBConnection, order_id); errOrder != nil {
		result = gin.H{
			"message": "Order not found!",
		}
	} else {
		dtrans, err := repository.GetDtransFromGivenId(database.DBConnection, order_id)

		if err != nil {
			panic(err)
		}

		result = gin.H{
			"order":        orders,
			"detail_order": dtrans,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetOrderFromAllUser(c *gin.Context) {
	orders, err := repository.GetAllOrders(database.DBConnection)

	var result gin.H
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		result = gin.H{
			"data": orders,
		}
	}

	c.JSON(http.StatusOK, result)
}

func CreateNewOrder(c *gin.Context) {
	newOrderID := strings.Replace(time.Now().Local().Format("20060102 150405"), " ", "", 1)
	err := repository.CreateNewOrder(database.DBConnection, newOrderID)

	var result string
	if err != nil {
		result = err.Error()
	} else {
		result = fmt.Sprintf("New order was successfully created #%s", newOrderID)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func GetDtransFromCurrentUser(c *gin.Context) {
	username, _, _ := c.Request.BasicAuth()
	dtrans, err := repository.GetDtransFromGivenId(database.DBConnection, username)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dtrans,
	})
}
