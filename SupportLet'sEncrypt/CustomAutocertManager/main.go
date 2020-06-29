package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context){
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:          autocert.AcceptTOS,
		Cache:           autocert.DirCache("/var/www/.cache"),
		HostPolicy:      autocert.HostWhitelist("example1.com", "example2.com"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
