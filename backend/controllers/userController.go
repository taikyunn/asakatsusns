package controller

import (
	"app/forms"
	entity "app/models/entity"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	db "app/models/db"
)

func SignUp(c *gin.Context) {

	// リクエスト内容を構造体に格納する
	form := forms.ValidateUser{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	// バリデーションの検証を行う
	if ok, errorMessages := form.Validate(); !ok {
		c.JSON(400, errorMessages)
		c.Abort()
		return
	}

	Name := c.PostForm("name")
	Email := c.PostForm("email")
	password := []byte(c.PostForm("password"))

	// パスワードをハッシュ化する。
	hashed, _ := bcrypt.GenerateFromPassword(password, 10)

	// `CompareHashAndPassword` でパスワードを検証する。
	err := bcrypt.CompareHashAndPassword(hashed, password)
	fmt.Println(err)

	// パスワードが間違っている場合は `bcrypt.ErrMismatchedHashAndPassword` を返す。
	err = bcrypt.CompareHashAndPassword(hashed, []byte("INCORRECT_PASSWORD"))
	fmt.Println(err == bcrypt.ErrMismatchedHashAndPassword)

	hashedPassword := string(hashed)

	var user = entity.User{
		Name:     Name,
		Email:    Email,
		Password: hashedPassword,
	}

	db.InsertUser(&user)
	c.JSON(200, user)
}

func GetAllUsers(c *gin.Context) {
	users := db.GetUserData()
	c.JSON(200, users)
}
