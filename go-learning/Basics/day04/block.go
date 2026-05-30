package main

import (
	
	"github.com/gin-gonic/gin"
)

// Middleware
func BlockMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(403, gin.H{
			"error": "Access blocked",
		})
		c.Abort()
	}
}
func main() {

	r := gin.Default()

	r.Use(BlockMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "srivignesh")
	})

	r.Run()
}