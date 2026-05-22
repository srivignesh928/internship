package main

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/glebarez/sqlite"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("DB connected")

	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Vignesh", Age: 21})

	fmt.Println("User inserted")
}