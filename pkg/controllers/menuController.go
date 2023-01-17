package controllers

import (
	"net/http"
	"strconv"

	"sanbercode-golang-final-project-nando/pkg/database"
	"sanbercode-golang-final-project-nando/pkg/repository"
	"sanbercode-golang-final-project-nando/pkg/structs"

	"github.com/gin-gonic/gin"
)

// MENU
func SearchMenu(c *gin.Context) {
	menu, err := repository.SearchMenu(database.DBConnection, c.Query(("menu_name")))

	var result gin.H
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		result = gin.H{
			"data": menu,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetAllMenu(c *gin.Context) {
	menu, err := repository.GetAllMenu(database.DBConnection)

	var result gin.H
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		result = gin.H{
			"data": menu,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetAllMenuWithTrashed(c *gin.Context) {
	menu, err := repository.GetAllMenuWithTrash(database.DBConnection)

	var result gin.H
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		result = gin.H{
			"data": menu,
		}
	}

	c.JSON(http.StatusOK, result)
}

func CreateNewMenu(c *gin.Context) {
	var menu structs.Menu
	var result string

	err := c.ShouldBindJSON(&menu)
	if err != nil {
		panic(err)
	}

	if menu.MenuPrice < 1000 {
		result = "Price must be greater than or equals Rp 1.000"
	} else {
		err = repository.CreateNewMenu(database.DBConnection, menu)
		if err != nil {
			result = err.Error()
		} else {
			result = "Success Insert Menu"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func UpdateMenu(c *gin.Context) {
	menu_id, err := strconv.Atoi(c.Query(("menu_id")))
	if err != nil {
		panic(err)
	}

	var result string
	if menu, errMenu := repository.GetMenu(database.DBConnection, menu_id); errMenu != nil {
		result = "Menu not found!"
	} else {
		err := c.ShouldBindJSON(&menu)
		if err != nil {
			panic(err)
		}

		if menu.MenuPrice < 0 {
			result = "Price must be greater than 0"
		} else {
			err = repository.UpdateMenu(database.DBConnection, menu)
			if err != nil {
				result = err.Error()
			} else {
				result = "Success Update Menu"
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func DeleteMenu(c *gin.Context) {
	menu_id, err := strconv.Atoi(c.Query(("menu_id")))
	if err != nil {
		panic(err)
	}

	var result string
	if menu, errMenu := repository.GetMenuWithTrashed(database.DBConnection, menu_id); errMenu != nil {
		result = "Menu not found!"
	} else {
		err = repository.DeleteMenu(database.DBConnection, menu)

		if err != nil {
			result = err.Error()
		} else {
			if menu.Deleted_at == nil {
				result = "Success Delete Menu"
			} else {
				result = "Success Restore Menu"
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
