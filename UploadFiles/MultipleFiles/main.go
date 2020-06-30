package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 Mib)
	router.MaxMultipartMemory= 8 << 20  // 8MiB
	router.POST("/upload", func(c *gin.Context){
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files  {
			dst := filepath.Base(file.Filename)
			log.Println(file.Filename)

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}
