package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "day10/docs"
)

// @title Day10 Swagger API
// @version 1.0
// @description Learning Swagger Documentation in Go.
// @host localhost:8080
// @BasePath /

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Srivignesh"},
	{ID: 2, Name: "Alex"},
}

func main() {

	r := gin.Default()

	r.GET("/hello", Hello)
	r.GET("/users", GetUsers)
	r.POST("/user", CreateUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// Hello godoc
//
// @Summary Get Hello Message
// @Description Returns a hello message
// @Tags Example
// @Produce json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Swagger!",
	})
}

// GetUsers godoc
//
// @Summary Get all users
// @Description Returns all users
// @Tags Users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
//
// @Summary Create a user
// @Description Adds a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body User true "User Data"
// @Success 201 {object} User
// @Router /user [post]
func CreateUser(c *gin.Context) {

	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input",
		})
		return
	}

	users = append(users, newUser)

	c.JSON(http.StatusCreated, newUser)
}