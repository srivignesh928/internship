package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Middleware
func simpleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before request")

		c.Next()

		fmt.Println("After request")
	}
}

func main() {

	r := gin.Default()

	r.Use(simpleMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "srivignesh")
	})

	r.Run()
}