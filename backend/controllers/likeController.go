package controller

import (
	db "app/models/db"
	"app/models/entity"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// いいね登録
func RegisterLikes(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)

	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	var likes = entity.Likes{
		ArticleId: articleID,
		UserId:    userID,
	}

	db.RegisterLikes(&likes)
}

// いいね解除
func DeleteLikes(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)

	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	db.DeleteLikes(articleID, userID)
}

// いいね数の取得
func GetCountFavorites(c *gin.Context) {
	// 直近10件の投稿データのarticleIdを取得
	articleIds := db.GetLastTenArticleID()
	log.Println("articleIds", articleIds)

	articleID := make([]int, len(articleIds))
	for i, v := range articleIds {
		articleID[i] = int(v.ID)
	}

	// article_idを元にcountデータを取得
	countData := db.GetLikeCount(articleID)
	c.JSON(200, countData)
}
