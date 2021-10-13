package main

import (
	"net/http"
	apicontroller "projectDD/apiController"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/parseJson", func(c *gin.Context) {
			apicontroller.JsonParser(c)
		})

		apiV1.POST("/parseJson", func(c *gin.Context) {
			apicontroller.JsonParser(c)
		})

	}

	router.Run(":8085") 
}
