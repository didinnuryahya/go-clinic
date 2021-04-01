package controller

import (
	"fmt"
	"go-clinic/database"
	"go-clinic/models"

	"github.com/gin-gonic/gin"
)

var tx = database.GetDB()

// GetUsers get all the user record form db
func GetPasiens(c *gin.Context) {
	var users []models.Pasien
	if err := tx.Find(&pasiens).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pasiens)
	}
}

// GetUser get single user record form db
func GetPasien(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func 
