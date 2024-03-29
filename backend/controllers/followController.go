package controller

import (
	db "app/models/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowData struct {
	FollowerCount int
	FollowCount   int
}

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

// フォロー解除
func DeleteFollow(c *gin.Context) {
	followerIdStr := c.PostForm("follower_id")
	followerId, _ := strconv.Atoi(followerIdStr)
	followedIdStr := c.PostForm("followed_id")
	followedId, _ := strconv.Atoi(followedIdStr)

	db.DeleteFollow(followerId, followedId)
}

// フォロワー情報取得
func GetFollowData(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userId, _ := strconv.Atoi(userIdStr)

	// フォロワー数の取得
	followerCount := db.GetFollower(userId)

	// フォロー数の取得
	followCount := db.GetFollow(userId)

	followData := []*FollowData{}
	followData = append(followData, &FollowData{followerCount, followCount})

	c.JSON(200, followData)
}

// フォロー一覧取得
func GetFollow(c *gin.Context) {
	followerIdStr := c.PostForm("follower_id")
	followerId, _ := strconv.Atoi(followerIdStr)

	// フォローしているユーザーIDを取得
	followedIds := db.GetFollowerList(followerId)

	// フォローしているユーザーの名前リストを取得
	followList := db.GetFollowNameList(followedIds)

	c.JSON(200, followList)
}

// フォロワー一覧取得
func GetFollower(c *gin.Context) {
	followedIdStr := c.PostForm("followed_id")
	followedId, _ := strconv.Atoi(followedIdStr)

	// フォロワーユーザーIDを取得
	followerIds := db.GetFollowedList(followedId)

	// フォロワーの名前リストの取得
	followerList := db.GetFollowNameList(followerIds)

	c.JSON(200, followerList)
}
