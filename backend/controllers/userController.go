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

	form := forms.ValidateUser{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	// バリデーション
	if ok, errorMessages := form.Validate(); !ok {
		c.JSON(201, errorMessages)
		return
	}

	Name := c.PostForm("name")
	Email := c.PostForm("email")
	password := c.PostForm("password")
	getHashedPassword(&password)

	var user = entity.User{
		Name:     Name,
		Email:    Email,
		Password: password,
	}

	db.InsertUser(&user)
	c.JSON(200, user)
}

func SignIn(c *gin.Context) {
	form := forms.LoginValidateUser{
		Name:     c.PostForm("name"),
		Password: c.PostForm("password"),
	}
	// バリデーション
	if ok, errorMessages := form.LoginValidate(); !ok {
		c.JSON(201, errorMessages)
		return
	}
	name := c.PostForm("name")
	password := c.PostForm("password")
	getHashedPassword(&password)

	// ユーザーレコードを取得する
	row := db.GetUserPassword(&name)
	fmt.Println(row)

}

func GetAllUsers(c *gin.Context) {
	users := db.GetUserData()
	c.JSON(200, users)
}

// パスワードのhash化を行う
func getHashedPassword(password *string) {
	Password := []byte(*password)
	// パスワードをハッシュ化する。
	hashed, _ := bcrypt.GenerateFromPassword(Password, 10)

	// `CompareHashAndPassword` でパスワードを検証する。
	err := bcrypt.CompareHashAndPassword(hashed, Password)
	fmt.Println(err)

	// パスワードが間違っている場合は `bcrypt.ErrMismatchedHashAndPassword` を返す。
	err = bcrypt.CompareHashAndPassword(hashed, []byte("INCORRECT_PASSWORD"))
	fmt.Println(err == bcrypt.ErrMismatchedHashAndPassword)

	*password = string(hashed)
}
