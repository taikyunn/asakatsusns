package db

import (
	entity "app/models/entity"
)

type TagInfo struct {
	ArticleId int
	TagId     []uint
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
		if err := db.Select("id,name").Where("id IN (?)", tagIDs).Find(&tag).Error; err != nil {
			panic(err.Error())
		}
		tagNames := make([]string, len(tag))
		tagIds := make([]uint, len(tag))
		for i, v := range tag {
			tagNames[i] = v.Name
			tagIds[i] = v.ID
		}
		tagInfo = append(tagInfo, &TagInfo{int(value), tagIds, tagNames})
	}
	return tagInfo
}

// タグ削除
func DeleteTags(articleId int) {
	db := gormConnect()
	var articleTag []entity.ArticleTag

	db.Delete(&articleTag, articleId)
	defer db.Close()
}

// タグ情報の取得
func FindTagData(articleId int) []string {
	db := gormConnect()
	var count int64
	var articleTag []entity.ArticleTag

	if err := db.Model(&articleTag).Where("article_id = ?", articleId).Count(&count).Error; err != nil {
		panic(err.Error())
	}

	if count == 0 {
		var null []string
		return null
	}

	// tag_idの取得
	if err := db.Select("tag_id").Where("article_id = ?", articleId).Find(&articleTag).Error; err != nil {
		panic(err.Error())
	}
	tagIDs := make([]uint, len(articleTag))
	for i, v := range articleTag {
		tagIDs[i] = uint(v.TagId)
	}

	// タグの中身を取得
	var tag []entity.Tag
	if err := db.Select("name").Where("id IN (?)", tagIDs).Find(&tag).Error; err != nil {
		panic(err.Error())
	}
	tagNames := make([]string, len(tag))
	for i, v := range tag {
		tagNames[i] = v.Name
	}
	return tagNames
}

// タグデータの更新
func UpdateTagData(articleId int, tags []entity.TagData) {
	db := gormConnect()
	var count int64
	var articleTag []entity.ArticleTag

	if err := db.Model(&articleTag).Where("article_id = ?", articleId).Count(&count).Error; err != nil {
		panic(err.Error())
	}

	if count != 0 {
		// 記事とタグの既存の紐付けを全解除＝deleted_atにデータを入れる
		db.Where("article_id = ?", articleId).Delete(&articleTag)
	}

	// tagテーブルにレコードか存在するか確認
	var tag []entity.Tag
	var tagId uint

	for _, value := range tags {
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

		// article_tagテーブルに登録
		var articleTag = entity.ArticleTag{
			ArticleId: int(articleId),
			TagId:     int(tagId),
		}
		db.Create(&articleTag)
	}
}

// 自動補完機能データ取得
func GetAutocompleteItems() []entity.Tag {
	db := gormConnect()
	var tag []entity.Tag

	if err := db.Select("name").Find(&tag).Error; err != nil {
		panic(err.Error())
	}
	return tag
}

// 投稿(1件)に紐づくタグデータを取得
func GetOneTagData(ArticleId int) []string {
	db := gormConnect()
	var count int
	var articleTag []entity.ArticleTag
	var tag []entity.Tag

	// article_tagテーブルにレコードが存在するか確認
	if err := db.Model(&articleTag).Where("article_id = ?", ArticleId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	if count == 0 {
		var null []string
		return null
	}

	// tag_idの取得
	if err := db.Select("tag_id").Where("article_id = ?", ArticleId).Find(&articleTag).Error; err != nil {
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
	return tagNames
}

// メインタグ情報の取得
func GetMainTag() []entity.Tag {
	db := gormConnect()
	var tag []entity.Tag

	if err := db.Select("id,name").Limit(5).Order("id DESC").Find(&tag).Error; err != nil {
		panic(err.Error())
	}
	return tag
}

// tagIdをもとにarticleIdを取得
func GetArticleIdByTagId(tagId int) []uint {
	db := gormConnect()
	var atricleTag []entity.ArticleTag

	if err := db.Select("article_id").Where("tag_id = ?", tagId).Limit(10).Order("id DESC").Find(&atricleTag).Error; err != nil {
		panic(err.Error())
	}
	articleIds := make([]uint, len(atricleTag))
	for i, v := range atricleTag {
		articleIds[i] = uint(v.ArticleId)
	}

	return articleIds
}

// Idに紐づいた中身を取得
func GetTagNameById(tagId int) []entity.Tag {
	db := gormConnect()
	var tag []entity.Tag

	if err := db.Select("name").Where("id = ?", tagId).Limit(10).Find(&tag).Error; err != nil {
		panic(err.Error())
	}
	return tag
}

// Idに紐づいたタグの件数を取得
func GetTagCount(tagId int) int {
	db := gormConnect()
	var articleTag []entity.ArticleTag
	var count int

	if err := db.Model(&articleTag).Where("tag_id = ?", tagId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	return count
}
