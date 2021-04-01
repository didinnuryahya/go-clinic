package controller

import (
	"fmt"
	"go-clinic/database"
	"go-clinic/models"
	"github.com/gin-gonic/gin"
)

var tx = database.GetDB()
var AdminModel []models.Admin
// GetUsers get all the user record form db
func GetAdmin(c *gin.Context) {
	var admin []models.Admin
	if err := tx.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

// GetUser get single user record form db
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.Admin
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := tx
	db.AutoMigrate([]models.Admin)
}

// What's bcrypt? https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
// You can change the value in bcrypt.DefaultCost to adjust the security index.
// 	err := userModel.setPassword("password0")
func (u *AdminModel) setPasswordAdmin(password string) error {
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
func (u *AdminModel) checkPasswordAdmin(password string) error {
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
func (model *AdminModel) UpdateAdmin(data interface{}) error {
	db := tx
	err := tx.Model(model).Update(data).Error
	return err
}


