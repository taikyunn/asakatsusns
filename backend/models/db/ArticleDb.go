package db

import (
	entity "app/models/entity"
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

	if err := db.Find(&articles).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return articles
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

// 編集
func UpdateArticleData(id int, body string) {
	db := gormConnect()
	var article []entity.Article

	db.Model(&article).Where("id = ?", id).Update("body", body)
	defer db.Close()
}
