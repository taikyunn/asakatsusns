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
