package controller

import (
	db "app/models/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckFollow(c *gin.Context) {
	followerIdStr := c.PostForm("follower_id")
	followerId, _ := strconv.Atoi(followerIdStr)

	isFollowed := db.CheckFollow(followerId)
	c.JSON(200, isFollowed)
}
