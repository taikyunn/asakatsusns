package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	entity "app/models/entity"
)

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

func InsertUser(signUp *entity.User) {
	db := gormConnect()

	db.Create(&signUp)
	defer db.Close()
}

func GetUserData() []entity.User {
	db := gormConnect()
	var user []entity.User

	db.Select("id,name").Find(&user)
	defer db.Close()
	return user
}

func GetUserPassword(name *string) []entity.User {
	db := gormConnect()
	var user []entity.User

	rows := db.Select("password").Where("name = ?", *name).Find(&user)
	fmt.Println(rows)
	defer db.Close()
	return user
}
