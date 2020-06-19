package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name	string	`form:"name"`
	Address	string	`form:"address"`
	Birthday	time.Time	`form:"birthday" time_format:"2006-01-02" tie_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}