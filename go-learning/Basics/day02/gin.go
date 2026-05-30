package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "i'm srivignesh,currently doing an internship at dotworld technologies",
		})
	})

	r.Run() // runs on localhost:8080
}