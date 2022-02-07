package db

import (
	entity "app/models/entity"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func gormConnect() *gorm.DB {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("ここでエラーが発生。")
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
	db.AutoMigrate(&entity.AchievementDay{})
	db.AutoMigrate(&entity.Article{})
	db.AutoMigrate(&entity.ArticleTag{})
	db.AutoMigrate(&entity.Comment{})
	db.AutoMigrate(&entity.Follow{})
	db.AutoMigrate(&entity.Likes{})
	db.AutoMigrate(&entity.Tag{})

	return db
}
