package controller

import (
	db "app/models/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

// フォローしているか判別
func CheckFollow(c *gin.Context) {
	followerIdStr := c.PostForm("follower_id")
	followerId, _ := strconv.Atoi(followerIdStr)
	followedIdStr := c.PostForm("followed_id")
	followedId, _ := strconv.Atoi(followedIdStr)

	isFollowed := db.CheckFollow(followerId, followedId)
	c.JSON(200, isFollowed)
}

// フォロー登録
func RegisterFollow(c *gin.Context) {
	followerIdStr := c.PostForm("follower_id")
	followerId, _ := strconv.Atoi(followerIdStr)
	followedIdStr := c.PostForm("followed_id")
	followedId, _ := strconv.Atoi(followedIdStr)

	db.RegisterFollow(followerId, followedId)
}
