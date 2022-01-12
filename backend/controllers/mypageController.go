package controller

import (
	db "app/models/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

// マイページ情報取得
func GetUserData(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)
	user := db.GetOneUser(userID)
	c.JSON(200, user)
}

// 寝る時間登録
func RegisterSleepTime(c *gin.Context) {
	sleepTime := c.PostForm("sleepTime")
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)
	db.UpdateSleepTime(sleepTime, userID)
}

// 起きる時間登録
func RegisterWakeUpTime(c *gin.Context) {
	wakeUpTime := c.PostForm("wakeUpTime")
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)
	db.UpdateWakeUpTime(wakeUpTime, userID)
}

// 寝る時間を取得
func GetSleepTimeData(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	user := db.GetSleepTimeById(userID)
	c.JSON(200, user)
}

// 寝る時間を編集
func UpdateSleepTime(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)
	sleepTime := c.PostForm("sleepTime")
	db.UpdateSleepTime(sleepTime, userID)
}

// 起きる時間取得
func GetWakeUpTimeData(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	user := db.GetWakeUpTimeById(userID)
	c.JSON(200, user)
}

// 起きる時間編集
func UpdateWakeUpTime(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)
	wakeUpTime := c.PostForm("wakeUpTime")
	db.UpdateWakeUpTime(wakeUpTime, userID)
}
