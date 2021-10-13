package apicontroller

import (
	"projectDD/structs"

	"github.com/gin-gonic/gin"
)

func JsonParser(c *gin.Context) {
	var Users structs.Users 

	c.BindJSON(&Users)

	c.HTML(200, "return.html", gin.H{
		"users": Users,
	})
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
