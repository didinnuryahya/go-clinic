package main

import (
	"go-clinic/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/admin/", admincontroller.GetAdmins)
		v1.GET("/admin/:id", admincontroller.GetAdmin)

		v2.GET("/apoteker/", apotekercontroller.GetApotekers)
		v2.GET("/apoteker/:id", apotekercontroller.GetApoteker)

		v3.GET("/dokter/", doktercontroller.GetDokters)
		v3.GET("/dokter/:id", controller.GetUser)

		v3.GET("/pasien/", controller.GetUsers)
		v3.GET("/pasien/:id", controller.GetUser)

	}
	v2 := router.Group("/api/v2")
	{
		
	}
	v3 := router.Group("/api/v3")
	{
	
	}
	v3 := router.Group("/api/v4")
	{
		
	}

	router.Run(":8080")
}
