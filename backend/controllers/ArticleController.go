package controller

import (
	db "app/models/db"
	entity "app/models/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	body := c.PostForm("body")
	name := c.PostForm("name")

	// user_idの取得
	dbUserId := db.GetUserIdByName(name)
	UserId := dbUserId[0].ID

	// データの登録
	var article = entity.Article{
		UserId: UserId,
		Body:   body,
	}

	db.InsertArticle(&article)
	c.JSON(200, gin.H{"mesage": "clear"})
}

// 投稿全件取得
func GetAllArticles(c *gin.Context) {
	articles := db.GetALLArticle()
	c.JSON(200, articles)
}

// 投稿削除
func DeleteArticle(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)
	db.DeleteArticleById(articleID)
}

// 編集画面・一件取得
func GetOneArticle(c *gin.Context) {
	articleIdStr := c.PostForm("id")
	articleID, _ := strconv.Atoi(articleIdStr)
	resultArticle := db.FindArticleData(articleID)
	c.JSON(200, resultArticle)
}

// 編集
func UpdateArticle(c *gin.Context) {
	articleIdStr := c.PostForm("id")
	articleID, _ := strconv.Atoi(articleIdStr)
	body := c.PostForm("body")
	db.UpdateArticleData(articleID, body)
}
