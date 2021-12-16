package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	controller "app/controllers"
)

func main() {
	// サーバーを起動する
	serve()
}

func serve() {
	r := gin.Default()

	// cors対策
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "https://api.asakatsusns.com"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 会員登録
	r.POST("/registerUser", controller.RegisterUser)

	r.Run(":3000")
}
