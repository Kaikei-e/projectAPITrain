package apicontroller

import (
	"projectDD/structs"

	"github.com/gin-gonic/gin"
)

func JsonParser(c *gin.Context) {
	var users structs.Users 

	c.BindJSON(&users)

	c.HTML(200, "return.html", gin.H{
		"users": users.Usrs,
	})
}
