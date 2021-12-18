package controller

import (
	entity "app/models/entity"

	"github.com/gin-gonic/gin"

	db "app/models/db"
)

func RegisterUser(c *gin.Context) {
	Name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user = entity.User{
		Name:     Name,
		Email:    email,
		Password: password,
	}
	db.InsertUser(&user)
	c.JSON(200, user)
}

func GetAllUsers(c *gin.Context) {
	users := db.GetUserData()
	c.JSON(200, users)
}
