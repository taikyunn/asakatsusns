package controller

import (
	"app/forms"
	entity "app/models/entity"
	"log"

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
	formPassword := c.PostForm("password")

	// パスワードの検証
	dbPassword := db.GetUserPassword(name).Password

	// ユーザーパスワードの比較
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword)); err != nil {
		log.Println("ログインできませんでした")
		c.JSON(201, gin.H{"err": err})
		c.Abort()
	} else {
		log.Println("ログインできました")
		c.JSON(200, name)
	}
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
