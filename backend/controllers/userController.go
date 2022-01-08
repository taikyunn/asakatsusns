package controller

import (
	"app/forms"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	db "app/models/db"
	entity "app/models/entity"
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

func Login(c *gin.Context) {
	// バリデーション
	form := forms.LoginValidateUser{
		Name:     c.PostForm("name"),
		Password: c.PostForm("password"),
	}

	if ok, errorMessages := form.LoginValidate(); !ok {
		c.JSON(201, errorMessages)
		return
	}

	// 名前とパスワードが一致するか確認
	name := c.PostForm("name")
	email := c.PostForm("email")

	if result := db.CheckNameAndPassword(name, email); !result {
		c.JSON(201, gin.H{"dbError": "*お名前とパスワードが一致しません。"})
		return
	}

	c.JSON(200, name)
}

func GetAllUsers(c *gin.Context) {
	users := db.GetUserData()
	c.JSON(200, users)
}

// パスワードのhash化を行う
func getHashedPassword(password *string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	*password = string(hash)
}

func GetSample(c *gin.Context) {
	UserName, _ := c.Get("UserName") // ログインユーザの取得
	log.Println("UserNameセッションの中身：", UserName)
}

func GetLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"UserName":     "",
		"ErrorMessage": "",
	})
}

func Public(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello public!\n"})
}

func Private(c *gin.Context) {

}
