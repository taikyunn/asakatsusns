package db

import (
	entity "app/models/entity"
	"log"
)

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

	for _, value := range tags {
		// tagtableにレコードか存在するか確認
		if err := db.Model(&tag).Where("name = ?", value.Text).Count(&count).Error; err != nil {
			panic(err.Error())
		}
		var tag = entity.Tag{
			Name: value.Text,
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
			if err := db.Select("id").Where("name = ?", value.Text).Find(&tag).Error; err != nil {
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

func GetArticleID() []entity.Article {
	db := gormConnect()
	var article []entity.Article

	if err := db.Select("id").Last(&article).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return article
}

// 投稿全件取得
func GetALLArticle() []entity.Article {
	db := gormConnect()
	var articles []entity.Article

	if err := db.Limit(10).Order("id DESC").Find(&articles).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return articles
}

type TagInfo struct {
	ArticleId int
	Name      []string
}

// タグ情報取得
func GetTagInfo(articleID []uint) []*TagInfo {
	db := gormConnect()
	var count int64
	var articleTag []entity.ArticleTag
	var tag []entity.Tag
	tagInfo := []*TagInfo{}

	for _, value := range articleID {
		// article_tagテーブルにレコードが存在するか確認
		if err := db.Model(&articleTag).Where("article_id = ?", value).Count(&count).Error; err != nil {
			panic(err.Error())
		}
		// 含まれていなかったらスキップ
		if count == 0 {
			continue
		}

		// tag_idの取得
		if err := db.Select("tag_id").Where("article_id = ?", value).Find(&articleTag).Error; err != nil {
			panic(err.Error())
		}
		tagIDs := make([]uint, len(articleTag))
		for i, v := range articleTag {
			tagIDs[i] = uint(v.TagId)
		}

		// タグの中身を取得
		if err := db.Select("name").Where("id IN (?)", tagIDs).Find(&tag).Error; err != nil {
			panic(err.Error())
		}
		tagNames := make([]string, len(tag))
		for i, v := range tag {
			tagNames[i] = v.Name
		}
		tagInfo = append(tagInfo, &TagInfo{int(value), tagNames})
	}
	return tagInfo
}

// 投稿削除
func DeleteArticleById(articleId int) {
	db := gormConnect()
	var article []entity.Article

	db.Delete(&article, articleId)
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

// タグ情報の取得
func FindTagData(articleId int) string {
	db := gormConnect()
	var count int64
	var articleTag []entity.ArticleTag

	if err := db.Model(&articleTag).Where("article_id = ?", articleId).Count(&count).Error; err != nil {
		panic(err.Error())
	}

	if count == 0 {
		return ""
	}

	// tag_idの取得
	if err := db.Select("tag_id").Where("article_id = ?", articleId).Find(&articleTag).Error; err != nil {
		panic(err.Error())
	}
	tagId := articleTag[0].TagId
	log.Println("articleId", articleId)
	log.Println("tagId", tagId)

	// tagの中身を取得
	var tag []entity.Tag
	if err := db.Select("name").Where("id = ?", tagId).Find(&tag).Error; err != nil {
		panic(err.Error())
	}
	tagInfo := tag[0].Name
	return tagInfo
}

// 編集
func UpdateArticleData(id int, body string) {
	db := gormConnect()
	var article []entity.Article

	db.Model(&article).Where("id = ?", id).Update("body", body)
	defer db.Close()
}
