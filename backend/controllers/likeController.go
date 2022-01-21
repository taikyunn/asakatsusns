package controller

import (
	db "app/models/db"
	"app/models/entity"
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
