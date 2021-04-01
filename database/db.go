package database

import (
	"fmt"
	"go-clinic/models"
	"github.com/jinzhu/gorm"

	// Register mysql driver for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"	
)

var db *gorm.DB 
var err error

const (
	DBHost = "127.0.0.1"

	DBUser = "root"

	DBPassword = ""

	DBName = "go-clinic"
)

func init() {
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	DbUser, DbPassword, DbHost, DbName)

	db, err = connectDB(dbinfo)
	if err != nil {
		fmt.Println(err)
	} else {
		db.AutoMigrate(&models.Admin{},
					   &models.Pasien{},
					   &models.Dokter{},
					   &models.Apoteker{},					   
					)
	}

}

// connectDB open a connection
func connectDB(dataSourceName string) (*gorm.DB, error) {
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("failed to connect database")
		return nil, err
	}
	return db, nil
}

//GetDB return database
func GetDB() *gorm.DB {
	return db
}