package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serves unicode entities
	r.GET("/json", func(c *gin.Context){
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves listeral characters
	r.GET("/purejson", func(c *gin.Context){

	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
