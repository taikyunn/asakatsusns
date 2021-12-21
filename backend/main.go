package main

import (
	"net/http"
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
		AllowOrigins:     []string{"http://localhost:8080", "https://front.asakatsusns.com"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/getAllUsers", controller.GetAllUsers)

	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLGlob("./templates/index.html")
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 会員登録
	r.POST("/signUp", controller.SignUp)

	// ログイン
	r.POST("/signIn", controller.SignIn)

	r.Run(":3000")
}
