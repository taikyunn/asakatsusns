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

	// メインタグ除法取得
	r.GET("/getMainTag", controller.GetMainTag)

	// 直近10件分のデータを取得
	r.GET("/getAllArticles", controller.GetAllArticles)

	// 無限スクロールのデータ取得
	r.POST("/getNextArticles", controller.GetNextArticles)

	// 投稿削除
	r.POST("/deleteArticle", controller.DeleteArticle)

	// 投稿編集画面表示
	r.POST("/getOneArticle", controller.GetOneArticle)

	// 投稿編集
	r.POST("/updateArticle", controller.UpdateArticle)

	// 投稿取得(マイページ)
	r.POST("/getMypageArticle", controller.GetMypageArticle)

	// 画像アップロード
	r.POST("/fileUpload", controller.FileUpload)

	// ユーザー画像情報取得
	r.POST("/getUserProfile", controller.GetUserProfile)

	// ユーザー名編集
	r.POST("/editUserName", controller.EditUserName)

	// いいね数取得(トップページ)
	r.GET("/getCountFavorites", controller.GetCountFavorites)

	// いいね数取得(無限スクロール)
	r.POST("/getNextCountFavorites", controller.GetNextCountFavorites)

	// いいねしているかの確認(無限スクロール)
	r.POST("/checkNextFavorite", controller.CheckNextFavorite)

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

	// フォローしているか判別
	r.POST("/checkFollow", controller.CheckFollow)

	// フォロー登録
	r.POST("/registerFollow", controller.RegisterFollow)

	// フォロー解除
	r.POST("/deleteFollow", controller.DeleteFollow)

	// フォロー情報取得
	r.POST("/getFollowData", controller.GetFollowData)

	// マイページ・いいねしているか確認
	r.POST("/checkFavoriteMypage", controller.CheckFavoriteMypage)

	// マイページ・いいね数取得
	r.POST("/getCountFavoriteMypage", controller.GetCountFavoriteMypage)

	// マイページ・いいねした投稿を取得
	r.POST("/getLikedPost", controller.GetLikedPost)

	// フォロー一覧取得
	r.POST("/getFollow", controller.GetFollow)

	// フォロワー一覧取得
	r.POST("/getFollower", controller.GetFollower)

	// タグごとの投稿一覧取得
	r.POST("/getTagArticles", controller.GetTagArticles)

	// いいねしているか判定(いいねした投稿)
	r.POST("/checkFavoriteLikedPost", controller.CheckFavoriteLikedPost)

	//いいね数の取得(いいねした投稿)
	r.POST("/getCountFavoriteLikedPost", controller.GetCountFavoriteLikedPost)

	// ランキング取得
	r.GET("/getWakeUpRanking", controller.GetWakeUpRanking)

	r.Run(":3000")
}
