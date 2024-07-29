package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ready to shorten URLs like a good Gurl !",
		})
	})
	// fmt.Println("a good Gurl shortens URLs with ease !")\

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web server -\nThis renders Gurl sad :(\nError: %v", err))
	}
}
