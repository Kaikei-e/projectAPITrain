package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router :=	gin.Default()

	router.LoadHTMLGlob("templates/")

	
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently ,"/")
	})
	
	router.GET("/", func (c *gin.Context)  {
		c.HTML(http.StatusOK, "index.html", gin.H{

		})
	})
  
	router.Run(":8085") // listen and serve on 0.0.0.0:8085 (for windows "localhost:8085")
}