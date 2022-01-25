package controller

import (
	db "app/models/db"
	"app/models/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

// コメントを登録
func InsertComment(c *gin.Context) {
	comment := c.PostForm("comment")
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	var commentData = entity.Comment{
		UserId:    userID,
		ArticleId: articleID,
		Comment:   comment,
	}

	db.InsertComment(&commentData)
}

// コメント件数取得
func GetCountComments(c *gin.Context) {
	// 直近10件の投稿データのarticleIdを取得
	articleIds := db.GetLastTenArticleID()

	articleID := make([]int, len(articleIds))
	for i, v := range articleIds {
		articleID[i] = int(v.ID)
	}

	// コメント件数を取得
	commentCount := db.GetCommentCount(articleID)
	c.JSON(200, commentCount)
}
