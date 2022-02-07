package db

import (
	entity "app/models/entity"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// ユーザーデータ登録
func InsertUser(signUp *entity.User) {
	db := gormConnect()

	db.Create(&signUp)
	defer db.Close()
}

func gormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ここでエラーです。")
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

// 正しい名前とパスワードの組み合わせかDBでチェック
func CheckNameAndPassword(name string, email string) bool {
	db := gormConnect()
	var user []entity.User

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

// ユーザーIdからユーザー名を取得(単数)
func GetUserName(userID uint) []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Select("name").Where("id = ?", userID).Find(&user).Error; err != nil {
		panic(err.Error())
	}
	return user
}
