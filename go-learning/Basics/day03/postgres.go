package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Model
type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {

	// DB connection
	dsn := "host=localhost user=postgres password=sql dbname=mydb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect DB:", err)
		return
	}

	fmt.Println("DB Connected")

	// Create table
	db.AutoMigrate(&User{})

	// Gin setup
	r := gin.Default()

	// CREATE
	r.POST("/user", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		db.Create(&user)
		c.JSON(200, user)
	})

	// READ
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, users)
	})

	// UPDATE
	r.PUT("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		var user User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		var input User
		c.BindJSON(&input)

		user.Name = input.Name
		user.Age = input.Age

		db.Save(&user)
		c.JSON(200, user)
	})

	// DELETE
	r.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		db.Delete(&User{}, id)

		c.JSON(200, gin.H{
			"message": "User deleted",
		})
	})

	// START SERVER (ONLY ONCE, AT END)
	r.Run()
}