package productcontroller

import (
	"github.com/AlifasulaemanIPAT/go-retapi-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context){

	var users []models.Users
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func Show(c *gin.Context){
	
}

func Create(c *gin.Context){
	
}

func Update(c *gin.Context){
	
}

func Delete(c *gin.Context){
	
}