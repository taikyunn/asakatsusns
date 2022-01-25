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

	// マイページ情報取得
	r.POST("/getUserData", controller.GetUserData)

	// 寝る時間登録
	r.POST("/registerSleepTime", controller.RegisterSleepTime)

	// 起きる時間登録
	r.POST("/registerWakeUpTime", controller.RegisterWakeUpTime)

	// 寝る時間取得
	r.POST("/getSleepTimeData", controller.GetSleepTimeData)

	// 寝る時間編集
	r.POST("/updateSleepTime", controller.UpdateSleepTime)

	// 起きる時間取得
	r.POST("/getWakeUpTimeData", controller.GetWakeUpTimeData)

	// 起きる時間編集
	r.POST("/updateWakeUpTime", controller.UpdateWakeUpTime)

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

	// 自動補完データの取得
	r.GET("/getAutocompleteItems", controller.GetAutocompleteItems)

	// 認証を必要とするリクエスト
	auth := r.Group("/post")
	auth.Use(middleware.AuthMiddleware())
	{
		// 記事投稿機能
		auth.POST("/new", controller.CreateArticle)

		// いいね登録
		auth.POST("/registerLikes", controller.RegisterLikes)

		// いいね削除
		auth.POST("/deleteLikes", controller.DeleteLikes)

		// コメント追加
		auth.POST("/insertComment", controller.InsertComment)
	}

	// 投稿全件取得
	r.GET("/getAllArticles", controller.GetAllArticles)

	// 投稿削除
	r.POST("/deleteArticle", controller.DeleteArticle)

	// 編集画面表示
	r.POST("/getOneArticle", controller.GetOneArticle)

	// 編集
	r.POST("/updateArticle", controller.UpdateArticle)

	// 画像アップロード
	r.POST("/fileUpload", controller.FileUpload)

	// ユーザー画像情報取得
	r.POST("/getUserProfile", controller.GetUserProfile)

	// ユーザー名編集
	r.POST("/editUserName", controller.EditUserName)

	// いいね数取得(トップページ)
	r.GET("/getCountFavorites", controller.GetCountFavorites)

	// いいね数取得(詳細ページ)
	r.POST("/getOneCountFavorites", controller.GetOneCountFavorites)

	// ログイン中のユーザーがいいねしているか確認
	r.POST("/checkFavorite", controller.CheckFavorite)

	// 投稿詳細ページ
	r.POST("/getArticleDetail", controller.GetArticleDetail)

	// 詳細ページ・ログイン中のユーザーがいいねしているか確認
	r.POST("/checkFavoriteByArticleId", controller.CheckFavoriteByArticleId)

	// コメント件数取得
	r.GET("/getCountComments", controller.GetCountComments)

	r.Run(":3000")
}
