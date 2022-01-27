package controller

import (
	db "app/models/db"
	"app/models/entity"
	"bytes"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
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

// マイページ画像登録
func FileUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	var request entity.Request
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	files := form.File["file"]

	// 乱数作成
	guid := xid.New()

	for _, file := range files {
		db.UploadS3Bucket(file, guid.String()+file.Filename)
		filepath := "images/" + guid.String() + file.Filename
		todoIdStr := request.ID
		todoID, _ := strconv.Atoi(todoIdStr)
		// filepathをmysqlに保存する
		db.UpdateFilePath(todoID, filepath)
	}

	c.JSON(200, gin.H{"message": "clear"})
}

func GetUserProfile(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	// filepathの取得
	user := db.GetFilePathById(userID)
	// デフォルト画像設定
	if len(user[0].ProfileImagePath) == 0 {
		user[0].ProfileImagePath = "images/default.png"
	}

	// filepathよりS3のデータを取得
	image, extension := db.DownloadS3Bucket(user[0].ProfileImagePath)

	var imageBytes []byte
	if extension == "png" {
		img, err := png.Decode(image)
		if err != nil {
			log.Fatal(err)
		}
		image.Close()

		// データをbyte型に変換
		buffer := new(bytes.Buffer)
		if err := png.Encode(buffer, img); err != nil {
			log.Println("unable to encode image.")
		}
		imageBytes = buffer.Bytes()

	} else {
		img, err := jpeg.Decode(image)
		if err != nil {
			log.Fatal(err)
		}
		image.Close()

		// データをbyte型に変換
		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, img, nil); err != nil {
			log.Println("unable to encode image.")
		}
		imageBytes = buffer.Bytes()
	}

	c.Data(200, "image/png", imageBytes)
}

// ユーザー名の編集
func EditUserName(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)
	userName := c.PostForm("name")

	db.UpdateUserName(userID, userName)
}

// 投稿・いいね数取得(マイページ)
func GetMypageArticle(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	mypageArticle, _ := db.GetMypageArticle(userID)

	c.JSON(200, mypageArticle)
}

// いいねしているか判定(マイページ)
func CheckFavoriteMypage(c *gin.Context) {
	mypageUserIdStr := c.PostForm("mypageUserId")
	mypageUserID, _ := strconv.Atoi(mypageUserIdStr)
	visiterUserIdStr := c.PostForm("visiterUserId")
	visiterUserID, _ := strconv.Atoi(visiterUserIdStr)

	// 投稿を取得
	_, articleIds := db.GetMypageArticle(mypageUserID)

	// いいね済みか判別
	favoriteData := db.CheckFavorite(articleIds, visiterUserID)

	c.JSON(200, favoriteData)
}

// マイページ・いいね数取得
func GetCountFavoriteMypage(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	_, articleIds := db.GetMypageArticle(userID)

	// いいね数を取得
	countData := db.GetLikeCount(articleIds)
	c.JSON(200, countData)
}
