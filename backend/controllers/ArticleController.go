package controller

import (
	"app/forms"
	db "app/models/db"
	"app/models/entity"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Id        int
	Name      string
	Body      string
	UserId    int
	UpdatedAt string
}

type DbTagResult struct {
	ArticleId int
	Key       uint
	Value     string
}

type EditData struct {
	ArticleId int
	UserId    int
	Body      string
	Tags      []string
}

type DetailData struct {
	ArticleId        int
	UserId           int
	Name             string
	Body             string
	ProfileImagePath string
	Image            string
	UpdatedAt        string
	Tags             []string
	Comments         []*db.ResultCommentData
	Count            int
}

func CreateArticle(c *gin.Context) {
	formData, _ := c.MultipartForm()
	var userData entity.UserData
	e := c.Bind(&userData)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	}

	// tagデータを構造体に入れる
	f := formData.Value
	tags := f["tags"][0]
	b := []byte(tags)
	var tagData []entity.TagData
	json.Unmarshal(b, &tagData)

	length := len(tagData)
	correctData := make([]string, 5)
	if length > 5 {
		for i, v := range tagData {
			if i == 5 {
				break
			}
			correctData[i] = v.Text
		}
	}

	form := forms.TagVaridator{
		Tags: tags,
	}

	if ok, errorMessages := form.TagValidate(); !ok {
		c.JSON(201, errorMessages)
		return
	}

	// user_idの取得
	dbUserId := db.GetUserIdByName(userData.Name)
	UserId := dbUserId[0].ID

	// データの登録
	var article = entity.Article{
		UserId: UserId,
		Body:   userData.Body,
	}

	db.InsertArticle(&article)

	// article_idの取得
	dbArticleId := db.GetArticleID()
	ArticleID := dbArticleId[0].ID

	// tag,article_tagテーブルにデータを登録
	if length > 5 {
		db.RegisterTags(ArticleID, correctData)
	} else {
		db.InsertTags(ArticleID, tagData)
	}

	// 早起きチェック
	count := checkWakeUptime(UserId)

	if count != 0 {
		c.JSON(200, count)
	}
}

// 直近10件分のデータを取得
func GetAllArticles(c *gin.Context) {
	// 直近10件分の投稿データを取得
	articles := db.GetALLArticle()

	articleID := make([]int, len(articles))
	for i, v := range articles {
		articleID[i] = v.Id
	}

	// タグ情報の取得
	tagInfo := db.GetTagInfo(articleID)

	tagMap := make(map[uint]string, len(tagInfo))
	for _, v := range tagInfo {
		for i := 0; i < len(v.Name); i++ {
			tagMap[v.TagId[i]] = v.Name[i]
		}
	}

	dbTagResult := []*DbTagResult{}

	for _, v := range tagInfo {
		for i := 0; i < len(v.TagId); i++ {
			for key := range tagMap {
				if (v.TagId[i]) == key {
					dbTagResult = append(dbTagResult, &DbTagResult{v.ArticleId, key, tagMap[key]})
				}
			}
		}
	}
	c.JSON(200, gin.H{"article": articles, "tag": dbTagResult})
}

// 投稿削除
func DeleteArticle(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)
	db.DeleteArticleById(articleID)
	db.DeleteTags(articleID)
}

// 編集画面・一件取得
func GetOneArticle(c *gin.Context) {
	articleIdStr := c.PostForm("id")
	articleID, _ := strconv.Atoi(articleIdStr)
	resultArticle := db.FindArticleData(articleID)

	// タグ情報の取得
	tagInfo := db.FindTagData(articleID)

	editData := []*EditData{}
	editData = append(editData, &EditData{int(resultArticle[0].ID), int(resultArticle[0].UserId), resultArticle[0].Body, tagInfo})

	c.JSON(200, editData)
}

// 編集
func UpdateArticle(c *gin.Context) {
	formData, _ := c.MultipartForm()
	var editArticleData entity.EditArticleData

	e := c.Bind(&editArticleData)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	}

	f := formData.Value
	tags := f["tags"][0]
	b := []byte(tags)
	var tagData []entity.TagData
	json.Unmarshal(b, &tagData)

	// 記事データの更新
	db.UpdateArticleData(editArticleData.Id, editArticleData.Body)

	// タグデータの更新
	db.UpdateTagData(editArticleData.Id, tagData)
}

// 詳細ページ表示
func GetArticleDetail(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)

	// 記事の中身を取得
	articleData := db.GetArticleBody(articleID)

	// タグデータの取得
	tagData := db.GetOneTagData(articleID)

	// コメントデータ取得
	commentData := db.GetCommentData(articleID)

	// コメント件数を取得
	count := db.GetOneCommentCount(articleID)

	detailData := []*DetailData{}
	detailData = append(detailData, &DetailData{articleID, int(articleData[0].UserId), articleData[0].Name, articleData[0].Body, articleData[0].ProfileImagePath, "", articleData[0].UpdatedAt, tagData, commentData, count})

	c.JSON(200, detailData)
}

// メインタグ情報の取得
func GetMainTag(c *gin.Context) {
	mainTag := db.GetMainTag()

	c.JSON(200, mainTag)
}

// 無限スクロールのデータ取得
func GetNextArticles(c *gin.Context) {
	countStr := c.PostForm("count")
	count, _ := strconv.Atoi(countStr)
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	var nextArticles []*db.NextArticleResult
	var countData []*db.CountData
	var commentCount []*db.CommentCount
	var favoriteData []*db.Favoritedata
	var dbTagResult []*DbTagResult
	var articleID []int

	// 一番古い投稿のupdatedAtを取得
	updatedAt, result := db.GetUpdatedAt(count)

	// 次の10件分のデータを取得
	if result {
		nextArticles = db.GetNextArticles(updatedAt)
		articleID = db.GetNextArticleID(updatedAt)
		countData = db.GetLikeCount(articleID)
		commentCount = db.GetCommentCount(articleID)
		favoriteData = db.CheckFavorite(articleID, userID)

		// タグ情報の取得
		tagInfo := db.GetTagInfo(articleID)
		tagMap := make(map[uint]string, len(tagInfo))
		for _, v := range tagInfo {
			for i := 0; i < len(v.Name); i++ {
				tagMap[v.TagId[i]] = v.Name[i]
			}
		}
		dbTagResult = []*DbTagResult{}

		for _, v := range tagInfo {
			for i := 0; i < len(v.TagId); i++ {
				for key := range tagMap {
					if (v.TagId[i]) == key {
						dbTagResult = append(dbTagResult, &DbTagResult{v.ArticleId, key, tagMap[key]})
					}
				}
			}
		}
	} else {
		c.JSON(201, gin.H{"message": "データがありません"})
		c.Abort()
	}
	c.JSON(200, gin.H{"nextArticles": nextArticles, "countData": countData, "commentCount": commentCount, "favoriteData": favoriteData, "tagData": dbTagResult})
}

// 早起きチェック
func checkWakeUptime(userId uint) int {

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	nowJST := time.Now().In(jst)

	const layout = "2006-01-02 15:04"
	yymmddmmssStr := nowJST.Format(layout)
	yymmddmmss := strings.Split(yymmddmmssStr, " ")

	yymmddData := yymmddmmss[0]
	mmssData := yymmddmmss[1]

	yymmddStr := strings.Split(yymmddData, "-")
	mmssStr := strings.Split(mmssData, ":")

	yearStr := yymmddStr[0]
	monthStr := yymmddStr[1]
	dayStr := yymmddStr[2]

	year, _ := strconv.Atoi(yearStr)
	month, _ := strconv.Atoi(monthStr)
	day, _ := strconv.Atoi(dayStr)

	hourStr := mmssStr[0]
	minuteStr := mmssStr[1]

	hour, _ := strconv.Atoi(hourStr)
	minute, _ := strconv.Atoi(minuteStr)

	wakeUpData := db.GetWakeUpData(int(userId))
	if wakeUpData[0].WakeUpTime == "" {
		return 0
	}

	wakeUpTime := wakeUpData[0].WakeUpTime
	wakeUphhmm := strings.Split(wakeUpTime, ":")
	rangeOfSuccess := wakeUpData[0].RangeOfSuccess

	wakeUpHourStr := wakeUphhmm[0]
	wakeUpMinuteStr := wakeUphhmm[1]

	wakeUpHour, _ := strconv.Atoi(wakeUpHourStr)
	wakeUpMinute, _ := strconv.Atoi(wakeUpMinuteStr)

	startHour := wakeUpHour - rangeOfSuccess

	startTime := time.Date(year, time.Month(month), day, startHour, wakeUpMinute, 00, 0, time.Local)
	endTime := time.Date(year, time.Month(month), day, wakeUpHour, wakeUpMinute, 00, 0, time.Local)
	targetTime := time.Date(year, time.Month(month), day, hour, minute, 00, 0, time.Local)

	if startTime.Before(targetTime) && targetTime.Before(endTime) {
		count := db.RegisterAchievementDay(targetTime, userId)

		if count != 0 {
			return count
		}
	}
	return 0
}
