package main

import (
	controller "app/controllers"
	middleware "app/middlewares"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	// ヘッダー情報
	r.POST("/getHeader", controller.GetHeader)

	// ユーザー全件取得
	r.GET("/getAllUsers", controller.GetAllUsers)

	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLGlob("./templates/index.html")
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 会員登録
	r.POST("/signUp", controller.SignUp)

	// ログイン
	r.POST("/login", controller.Login)

	// 投稿ここで認証をしてから次に行けるように設定する。
	menu := r.Group("/post")
	menu.Use(middleware.AuthMiddleware())
	{
		menu.POST("/new", controller.Create)
	}

	// 投稿全件取得
	r.GET("/getAllArticles", controller.GetAllArticles)

	// 投稿削除
	r.POST("/deleteArticle", controller.DeleteArticle)

	// 編集画面表示
	r.POST("/getOneArticle", controller.GetOneArticle)

	// 編集
	r.POST("/updateArticle", controller.UpdateArticle)

	r.Run(":3000")
}
