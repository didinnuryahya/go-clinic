package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-clinic/database"
	"go-clinic/models"
)

var tx = database.GetDB()
var PembayaranModel []models.Pembayaran

// GetUsers get all the user record form db
func GetPembayarans(c *gin.Context) {
	var pembayarans []models.Pembayaran
	if err := tx.Find(&pembayarans).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pembayarans)
	}
}

// GetUser get single user record form db
func GetPembayaran(c *gin.Context) {
	id := c.Params.ByName("id")
	var pembayaran []models.Pembayaran
	if err := tx.Where("id = ?", id).First(&pembayaran).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pembayaran)
	}
}

// Migrate the schema of database if needed
func AutoMigratePembayaran() {
	db := tx
	db.AutoMigrate([]models.Pembayaran)
}

// You could input the conditions and it will return an UserModel in database with error info.
// 	userModel, err := FindOneUser(&UserModel{Username: "username0"})
func SaveOnePembayaran(data interface{}) error {
	db := tx
	err := tx.Save(data).Error
	return err
}

// You could update properties of an UserModel to database returning with error info.
//  err := db.Model(userModel).Update(UserModel{Username: "wangzitian0"}).Error
func (model *PembayaranModel) UpdatePembayaran(data interface{}) error {
	db := tx
	err := tx.Model(model).Update(data).Error
	return err
}

func DeletePembayaranModel(condition interface{}) error {
	db := database.GetDB()
	err := db.Where(condition).Delete(PembayaranModel{}).Error
	return err
}
