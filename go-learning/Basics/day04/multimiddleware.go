package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Middleware
func M1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("M1 start")
		c.Next()
		fmt.Println("M1 end")
	}
}

func M2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("M2 start")
		c.Next()
		fmt.Println("M2 end")
	}
}
func main() {

	r := gin.Default()

	r.Use(M1(), M2())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "srivignesh")
	})

	r.Run()
}