package apicontroller

import (
	//	"fmt"
	"projectDD/structs"

	"github.com/gin-gonic/gin"
)

func JsonParser(c *gin.Context) {
	var theUser structs.User 
	//var data []byte

	//dataByID := c.PostForm("jsData")
	//c.Data(400, "body", data)

	//dataS := string(data)

	//fmt.Println(dataS)
	//fmt.Println(dataByID)


	err := c.BindJSON(&theUser)
	if err != nil {
		panic(err)
	}

	/*
	c.HTML(200, "return.html", gin.H{
		"users": theUser,
	})
	*/


	c.JSON(200, gin.H{
		"status": "ok",
		"data": theUser,
	})
	
}
