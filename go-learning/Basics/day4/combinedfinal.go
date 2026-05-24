package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


//  Logger Middleware
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request started:", c.Request.URL.Path)

		c.Next()

		fmt.Println(" Response sent:", c.Request.URL.Path)
	}
}


//  Recovery Middleware
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(" Recovered from panic:", err)

				c.JSON(500, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		c.Next()
	}
}


func main() {

	r := gin.New() // important: no default middleware

	//  Add both middleware
	r.Use(LoggerMiddleware(), RecoveryMiddleware())

	// Normal route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello Srivignesh")
	})

	// Panic route
	r.GET("/panic", func(c *gin.Context) {
		panic("Something went wrong!")
	})

	r.Run()
}