package main

import (
	"github.com/AlifasulaemanIPAT/go-retapi-gin/models"
	"github.com/AlifasulaemanIPAT/go-retapi-gin/controllers/productcontroller"
	
	"github.com/gin-gonic/gin"

)

func main () {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/api/users", productcontroller.Index)
	r.GET("/api/users/:id", productcontroller.Show)
	r.POST("/api/users", productcontroller.Create)
	r.PUT("/api/users/:id", productcontroller.Update)
	r.DELETE("/api/users", productcontroller.Delete)

	r.Run()
}