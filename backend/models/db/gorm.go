package db

import (
	entity "app/models/entity"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func gormConnect() *gorm.DB {
	err := godotenv.Load()
	// todo環境変数が本番環境だと読み込めていない。
	if err != nil {
		DBMS := "mysql"
		DBNAME := "asakatsusns"
		USER := "admin"
		PASS := "password"
		ADDRESS := "asakatsusns.cvfat2usql6c.ap-northeast-1.rds.amazonaws.com"

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
	}

	DBMS := "mysql"
	DBNAME := os.Getenv("DB_NAME")
	USER := os.Getenv("API_USER")
	PASS := os.Getenv("API_PASS")
	ADDRESS := os.Getenv("API_ADDRESS")

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
