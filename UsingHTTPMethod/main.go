package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Router Get
func getting(c *gin.Context) {
	log.Println("Get")
}

// Router Post
func posting(c *gin.Context) {
	log.Println("Post")
}

// Router Put
func putting(c *gin.Context) {
	log.Println("Put")
}

// Router DELETE
func deleting(c *gin.Context) {
	log.Println("Delete")
}

// Router Patch
func patching(c  *gin.Context) {
	log.Println("Patch")
}

// Router Head
func head(c *gin.Context) {
	log.Println("Head")
}

// Router Options
func options(c *gin.Context) {
	log.Println("Options")
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default it serve on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
