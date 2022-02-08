package db

import (
	entity "app/models/entity"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func gormConnect() *gorm.DB {
	var USER string
	var PASS string
	var ADDRESS string
	var DBNAME string

	err := godotenv.Load("env/dev.env")
	if err != nil {
		USER = os.Getenv("DB_USER")
		PASS = os.Getenv("DB_PASS")
		ADDRESS = os.Getenv("DB_ADDRESS")
		DBNAME = os.Getenv("DB_NAME")
	} else {
		USER = os.Getenv("API_USER")
		PASS = os.Getenv("API_PASS")
		ADDRESS = os.Getenv("API_ADDRESS")
		DBNAME = os.Getenv("DB_NAME")
	}
	DBMS := "mysql"

	// コンテナ名:ポート番号を指定する
	CONNECT := USER + ":" + PASS + "@tcp(" + ADDRESS + ":3306)" + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	log.Println(CONNECT)

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)

	db.SingularTable(true)

	db.AutoMigrate(&entity.User{})

	log.Println("db connected ", &db)

	return db

}
