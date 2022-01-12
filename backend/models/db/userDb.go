package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	entity "app/models/entity"
)

// User モデルの宣言
type User struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}

func gormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	USER := os.Getenv("API_USER")
	PASS := os.Getenv("API_PASS")
	ADDRESS := os.Getenv("API_ADDRESS")
	DBMS := "mysql"
	DBNAME := os.Getenv("DB_NAME")

	if os.Getenv("DB_ENV") == "production" {
		USER = os.Getenv("DB_USER")
		PASS = os.Getenv("DB_PASS")
		ADDRESS = os.Getenv("DB_ADDRESS")
	}

	// コンテナ名:ポート番号を指定する
	CONNECT := USER + ":" + PASS + "@tcp(" + ADDRESS + ":3306)" + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	fmt.Println(CONNECT)

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)

	db.SingularTable(true)

	db.AutoMigrate(&entity.User{})

	fmt.Println("db connected ", &db)

	return db
}

// ユーザーデータ登録
func InsertUser(signUp *entity.User) {
	db := gormConnect()

	db.Create(&signUp)
	defer db.Close()
}

// ユーザーデータの取得
func GetUserData() []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Select("id,name").Find(&user).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return user
}

// 正しい名前とパスワードの組み合わせかDBでチェック
func CheckNameAndPassword(name string, email string) bool {
	db := gormConnect()
	var user User

	if err := db.Where("name = ? AND email = ?", name, email).Find(&user).Error; err != nil {
		return false
	}
	return true
}

// ユーザーIDの配列からユーザー名を取得
func GetNameById(userID []uint) []entity.User {
	db := gormConnect()
	var user []entity.User

	db.Select("id,name").Find(&user, userID)
	defer db.Close()
	return user
}
