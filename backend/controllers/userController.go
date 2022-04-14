package controller

import (
	"app/forms"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	db "app/models/db"
	entity "app/models/entity"
)

// ヘッダー情報取得
func GetHeader(c *gin.Context) {
	userIdStr := c.PostForm("id")
	userID, _ := strconv.Atoi(userIdStr)
	c.JSON(200, userID)
}

// ユーザー登録
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

	// nameのuniqueチェック
	if ok, errorMessages := db.CheckNameUnique(Name); !ok {
		c.JSON(201, errorMessages)
		return
	}
	getHashedPassword(&password)

	var user = entity.User{
		Name:           Name,
		Email:          Email,
		Password:       password,
		RangeOfSuccess: 3,
	}

	db.InsertUser(&user)
	c.JSON(200, user)
}

// ログイン処理
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

	// ユーザーIDを取得
	dbUserId := db.GetUserIdByName(name)

	c.JSON(200, gin.H{"name": name, "userId": dbUserId})
}

// パスワードのhash化を行う
func getHashedPassword(password *string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	*password = string(hash)
}

// ImagePathを取得
func GetLoginUserProfileImagePath(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	path := db.GetImagePath(userID)
	c.JSON(200, path)
}
