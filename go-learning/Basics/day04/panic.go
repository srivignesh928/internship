package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Recovery Middleware
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered:", err)

				c.JSON(500, gin.H{
					"error": "Something went wrong",
				})
			}
		}()

		c.Next()
	}
}

func main() {

	r := gin.Default()

	// Use middleware
	r.Use(RecoveryMiddleware())

	// Normal route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello")
	})

	// Panic route
	r.GET("/panic", func(c *gin.Context) {
		panic("Something broke!")
	})

	r.Run()
}