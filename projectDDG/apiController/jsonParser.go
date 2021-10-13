package apicontroller

import (
	"fmt"
	"log"
	"projectDD/structs"

	"github.com/gin-gonic/gin"
)

func JsonParser(c *gin.Context) {
	var Users structs.Users 

	fmt.Println(c.Params)

	err := c.BindJSON(&Users)
	if err != nil {
		log.Fatalln(err)
	}

	/*
	c.HTML(200, "return.html", gin.H{
		"users": Users,
	})
	*/
	c.JSON(200, gin.H{
		"status": "ok",
		"data": Users,
	})
	
}
