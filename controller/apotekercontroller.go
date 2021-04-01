package controller

import (
	"fmt"
	"go-clinic/database"
	"go-clinic/models"
	"github.com/gin-gonic/gin"
)

var tx = database.GetDB()
var AdminModel models.Apoteker
// GetUsers get all the user record form db
func GetApoteker(c *gin.Context) {
	var users []models.User
	if err := tx.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

// GetUser get single user record form db
func GetApoteker(c *gin.Context) {
	id := c.Params.ByName("id")
	var admin models.Admin
	if err := tx.Where("id = ?", id).First(&apoteker).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

// Migrate the schema of database if needed
func AutoMigrateApoteker() {
	db := tx
	db.AutoMigrate([]models.Apoteker)
}

// What's bcrypt? https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
// You can change the value in bcrypt.DefaultCost to adjust the security index.
// 	err := userModel.setPassword("password0")
func (u *AdminModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password tidak boleh kosong!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *AdminModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// You could input the conditions and it will return an UserModel in database with error info.
// 	userModel, err := FindOneUser(&UserModel{Username: "username0"})
func SaveOne(data interface{}) error {
	db := tx
	err := tx.Save(data).Error
	return err
}

// You could update properties of an UserModel to database returning with error info.
//  err := db.Model(userModel).Update(UserModel{Username: "wangzitian0"}).Error
func (model *AdminModel) Update(data interface{}) error {
	db := tx
	err := tx.Model(model).Update(data).Error
	return err
}


