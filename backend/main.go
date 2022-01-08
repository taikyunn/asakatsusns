package main

import (
	controller "app/controllers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/api/option"
)

func main() {
	// サーバーを起動する
	serve()
}

func serve() {
	r := gin.Default()

	// cors対策
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080", "https://front.asakatsusns.com"},
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
			"PUT",
		},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ユーザー全件取得
	r.GET("/getAllUsers", controller.GetAllUsers)

	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLGlob("./templates/index.html")
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 会員登録
	r.POST("/signUp", controller.SignUp)

	// ログイン
	r.POST("/login", controller.Login)

	r.GET("/public", controller.Public)

	// 投稿ここで認証をしてから次に行けるように設定する。
	menu := r.Group("/post")
	menu.Use(authMiddleware())
	{
		menu.POST("/new", controller.Create)
	}

	r.Run(":3000")
}

func authMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Firebase SDK のセットアップ
		opt := option.WithCredentialsFile("firebase-adminsdk.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		// クライアントから送られてきた JWT 取得
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			u := fmt.Sprintf("error verifying ID token: %v\n", err)
			c.JSON(http.StatusBadRequest, u)
			return
		}
		log.Printf("Verified ID token: %v\n", token)
		c.Next()
	})
}
