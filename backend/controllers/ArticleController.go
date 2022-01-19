package controller

import (
	"app/forms"
	db "app/models/db"
	"app/models/entity"
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

func CreateArticle(c *gin.Context) {
	formData, _ := c.MultipartForm()
	var userData entity.UserData
	e := c.Bind(&userData)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	}
	log.Println("names.Name:", userData.Name)

	// mapデータの取り出し方を確
	t := formData.Value
	tags := t["tags"][0]
	//tagデータ
	log.Println(tags)

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
	c.JSON(200, gin.H{"mesage": "clear"})
}

// 投稿全件取得
func GetAllArticles(c *gin.Context) {
	// 投稿を全件取得
	articles := db.GetALLArticle()

	userID := make([]uint, len(articles))
	for i, v := range articles {
		userID[i] = v.UserId
	}

	result := []*Result{}

	// idより投稿者を取得
	user := db.GetNameById(userID)

	// 返すデータの作成
	for _, av := range articles {
		for _, uv := range user {
			if av.UserId == uv.ID {
				result = append(result, &Result{int(av.ID), uv.Name, av.Body, int(av.UserId)})
			}
		}
	}

	c.JSON(200, result)
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
	c.JSON(200, resultArticle)
}

// 編集
func UpdateArticle(c *gin.Context) {
	articleIdStr := c.PostForm("id")
	articleID, _ := strconv.Atoi(articleIdStr)
	body := c.PostForm("body")
	db.UpdateArticleData(articleID, body)
}
