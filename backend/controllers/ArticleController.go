package controller

import (
	"app/forms"
	db "app/models/db"
	"app/models/entity"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Id     int
	Name   string
	Body   string
	UserId int
}

type DbTagResult struct {
	Id  int
	Tag []string
}

type EditData struct {
	ArticleId int
	UserId    int
	Body      string
	Tags      []string
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

	// バリデーション
	form := forms.ArticleValidate{
		Body: c.PostForm("body"),
	}

	if ok, errorMessages := form.ArticleValidate(); !ok {
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
	db.InsertTags(ArticleID, tagData)

	c.JSON(200, gin.H{"mesage": "clear"})
}

// 投稿全件取得
func GetAllArticles(c *gin.Context) {
	// 投稿を10件取得
	articles := db.GetALLArticle()

	userID := make([]uint, len(articles))
	articleID := make([]uint, len(articles))
	for i, v := range articles {
		userID[i] = v.UserId
		articleID[i] = v.ID
	}

	// タグ情報の取得
	tagInfo := db.GetTagInfo(articleID)

	dbTagResult := []*DbTagResult{}

	for _, v := range tagInfo {
		dbTagResult = append(dbTagResult, &DbTagResult{v.ArticleId, v.Name})
	}

	// idより投稿者を取得
	user := db.GetNameById(userID)

	result := []*Result{}

	// 返すデータの作成
	for _, av := range articles {
		for _, uv := range user {
			if av.UserId == uv.ID {
				result = append(result, &Result{int(av.ID), uv.Name, av.Body, int(av.UserId)})
			}
		}
	}

	c.JSON(200, gin.H{
		"article": result,
		"tag":     dbTagResult,
	})
}

// 投稿削除
func DeleteArticle(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)
	db.DeleteArticleById(articleID)
}

// 編集画面・一件取得
func GetOneArticle(c *gin.Context) {
	articleIdStr := c.PostForm("id")
	articleID, _ := strconv.Atoi(articleIdStr)
	resultArticle := db.FindArticleData(articleID)

	// タグ情報の取得
	tagInfo := db.FindTagData(articleID)
	log.Println("taginfo", tagInfo)

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
