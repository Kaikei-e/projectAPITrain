package apicontroller

import (
	//	"fmt"
	"log"
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
		log.Fatalln(err)
	}

	/*
	c.HTML(200, "return.html", gin.H{
		"users": Users,
	})
	*/
	c.JSON(200, gin.H{
		"status": "ok",
		"data": theUser,
	})
	
}
