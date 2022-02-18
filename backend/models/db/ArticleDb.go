package db

import (
	entity "app/models/entity"
	"time"
)

type NextArticle struct {
	Id               int
	UserId           int
	Body             string
	UpdatedAt        time.Time
	Name             string
	ProfileImagePath string
}

type NextArticleResult struct {
	Id               int
	UserId           int
	Body             string
	UpdatedAt        string
	Name             string
	ProfileImagePath string
	Image            string
}

type ArticleData struct {
	UserId           int
	Name             string
	Body             string
	ProfileImagePath string
	UpdatedAt        time.Time
}

type ArticleResult struct {
	UserId           int
	Name             string
	Body             string
	ProfileImagePath string
	UpdatedAt        string
}

// userIdの取得
func GetUserIdByName(name string) []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Select("id").Where("name = ?", name).Find(&user).Error; err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return user
}

// 投稿を保存
func InsertArticle(article *entity.Article) bool {
	db := gormConnect()

	db.Create(&article)
	defer db.Close()

	return true
}

func InsertTags(articleId uint, tags []entity.TagData) {
	db := gormConnect()
	var tag []entity.Tag
	var count int64
	var tagId uint

	for _, v := range tags {
		// tagテーブルにレコードか存在するか確認
		if err := db.Model(&tag).Where("name = ?", v.Text).Count(&count).Error; err != nil {
			panic(err.Error())
		}

		var tag = entity.Tag{
			Name: v.Text,
		}

		// tagテーブルにデータを登録
		if count == 0 {
			db.Create(&tag)

			// tag_idの取得
			if err := db.Select("id").Last(&tag).Error; err != nil {
				panic(err.Error())
			}
			tagId = tag.ID

		} else {
			// tag_idの取得
			if err := db.Select("id").Where("name = ?", v.Text).Find(&tag).Error; err != nil {
				panic(err.Error())
			}
			tagId = tag.ID
		}

		// article_tagテーブルにレコードを保存する
		var articleTag = entity.ArticleTag{
			ArticleId: int(articleId),
			TagId:     int(tagId),
		}
		db.Create(&articleTag)
	}
}

func RegisterTags(articleId uint, tags []string) {
	db := gormConnect()
	var tag []entity.Tag
	var count int64
	var tagId uint

	for _, v := range tags {
		// tagテーブルにレコードか存在するか確認
		if err := db.Model(&tag).Where("name = ?", v).Count(&count).Error; err != nil {
			panic(err.Error())
		}

		var tag = entity.Tag{
			Name: v,
		}

		// tagテーブルにデータを登録
		if count == 0 {
			db.Create(&tag)

			// tag_idの取得
			if err := db.Select("id").Last(&tag).Error; err != nil {
				panic(err.Error())
			}
			tagId = tag.ID

		} else {
			// tag_idの取得
			if err := db.Select("id").Where("name = ?", v).Find(&tag).Error; err != nil {
				panic(err.Error())
			}
			tagId = tag.ID
		}

		// article_tagテーブルにレコードを保存する
		var articleTag = entity.ArticleTag{
			ArticleId: int(articleId),
			TagId:     int(tagId),
		}
		db.Create(&articleTag)
	}
}

// 記事IDの取得
func GetArticleID() []entity.Article {
	db := gormConnect()
	var article []entity.Article

	if err := db.Select("id").Last(&article).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return article
}

// 直近10件の取得
func GetALLArticle() []*NextArticleResult {
	db := gormConnect()
	nextArticle := []*NextArticle{}
	nextArticleResult := []*NextArticleResult{}
	column := "article.id, user_id, body, article.updated_at, name, profile_image_path"
	table := "INNER JOIN user ON article.user_id = user.id"
	where := "article.deleted_at IS NULL AND user.deleted_at IS NULL"

	if err := db.Table("article").Select(column).Joins(table).Where(where).Limit(10).Order("updated_at DESC").Scan(&nextArticle).Error; err != nil {
		panic(err.Error())
	}
	for _, v := range nextArticle {
		t := v.UpdatedAt.Format("2006/01/02 15:04:05")
		nextArticleResult = append(nextArticleResult, &NextArticleResult{int(v.Id), v.UserId, v.Body, t, v.Name, v.ProfileImagePath, ""})
	}
	defer db.Close()

	return nextArticleResult
}

// 投稿削除
func DeleteArticleById(articleId int) {
	db := gormConnect()
	var article []entity.Article

	db.Where("id = ?", articleId).Delete(&article)
	defer db.Close()
}

// 投稿1件取得
func FindArticleData(articleId int) []entity.Article {
	db := gormConnect()
	var article []entity.Article

	db.First(&article, articleId)
	defer db.Close()
	return article
}

// 記事情報更新
func UpdateArticleData(id int, body string) {
	db := gormConnect()
	var article []entity.Article

	db.Model(&article).Where("id = ?", id).Update("body", body)
	defer db.Close()
}

// body,user_id, updated_atを取得
func GetArticleBody(articleID int) []*ArticleResult {
	db := gormConnect()
	articleData := []*ArticleData{}
	articleResult := []*ArticleResult{}
	table := "user_id, name, body, article.updated_at, profile_image_path"
	join := "INNER JOIN user ON article.user_id = user.id"
	where := "article.deleted_at IS NULL AND user.deleted_at IS NULL AND article.id = ?"

	if err := db.Table("article").Select(table).Joins(join).Where(where, articleID).Scan(&articleData).Error; err != nil {
		panic(err.Error())
	}
	for _, v := range articleData {
		t := v.UpdatedAt.Format("2006/01/02 15:04:05")
		articleResult = append(articleResult, &ArticleResult{int(v.UserId), v.Name, v.Body, v.ProfileImagePath, t})
	}
	defer db.Close()
	return articleResult
}

// 直近10件の取得
func GetArticleByTag(articleID []int) []*NextArticleResult {
	db := gormConnect()
	nextArticle := []*NextArticle{}
	nextArticleResult := []*NextArticleResult{}
	column := "article.id, user_id, body, article.updated_at, name, profile_image_path"
	table := "INNER JOIN user ON article.user_id = user.id"

	for _, v := range articleID {
		if err := db.Table("article").Select(column).Joins(table).Limit(10).Order("updated_at DESC").Where("article.deleted_at IS NULL AND user.deleted_at IS NULL AND article.id = ?", v).Scan(&nextArticle).Error; err != nil {
			panic(err.Error())
		}
		for _, v := range nextArticle {
			t := v.UpdatedAt.Format("2006/01/02 15:04:05")
			nextArticleResult = append(nextArticleResult, &NextArticleResult{int(v.Id), v.UserId, v.Body, t, v.Name, v.ProfileImagePath, ""})
		}
	}

	defer db.Close()

	return nextArticleResult
}

// 指定したupdated_atを取得
func GetUpdatedAt(count int) (time.Time, bool) {
	db := gormConnect()
	var articles []entity.Article
	count = count * 10
	var updatedAt time.Time
	var result bool

	if err := db.Select("updated_at").Order("updated_at DESC").Limit(count).Find(&articles).Error; err != nil {
		panic(err.Error())
	}

	length := len(articles) - 1
	if count-1 == length {
		updatedAt = articles[length].UpdatedAt
		result = true
	} else {
		updatedAt = time.Time{}
		result = false
	}
	defer db.Close()
	return updatedAt, result
}

// 次の10件分のデータを取得
func GetNextArticles(updatedAt time.Time) []*NextArticleResult {
	db := gormConnect()
	nextArticle := []*NextArticle{}
	nextArticleResult := []*NextArticleResult{}
	column := "article.id, user_id, body, article.updated_at, name, profile_image_path"
	table := "INNER JOIN user ON article.user_id = user.id"

	if err := db.Table("article").Select(column).Joins(table).Where("article.deleted_at IS NULL AND article.updated_at < ?", updatedAt).Limit(10).Order("updated_at DESC").Scan(&nextArticle).Error; err != nil {
		panic(err.Error())
	}
	for _, v := range nextArticle {
		t := v.UpdatedAt.Format("2006/01/02 15:04:05")
		nextArticleResult = append(nextArticleResult, &NextArticleResult{int(v.Id), v.UserId, v.Body, t, v.Name, v.ProfileImagePath, ""})
	}
	defer db.Close()
	return nextArticleResult
}
