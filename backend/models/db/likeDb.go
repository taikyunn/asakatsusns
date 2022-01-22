package db

import (
	entity "app/models/entity"
)

type CountData struct {
	ArticleId int
	Count     int
}

type Favoritedata struct {
	ArticleId int
	Count     bool
}

// いいね登録
func RegisterLikes(like *entity.Likes) {
	db := gormConnect()

	db.Create(&like)
	defer db.Close()
}

// いいね解除
func DeleteLikes(articleId int, userId int) {
	db := gormConnect()
	var likes []entity.Likes

	db.Where("user_id = ? AND article_id = ? ", userId, articleId).Delete(&likes)
	defer db.Close()
}

// 直近10件のarticleIdを取得
func GetLastTenArticleID() []entity.Article {
	db := gormConnect()
	var article []entity.Article

	if err := db.Select("id").Limit(10).Order("id DESC").Find(&article).Error; err != nil {
		panic(err.Error())
	}
	return article
}

// いいね数を取得する
func GetLikeCount(articleIds []int) []*CountData {
	db := gormConnect()
	var likes []entity.Likes
	var count int
	countData := []*CountData{}

	for _, v := range articleIds {
		db.Where("article_id = ?", v).Find(&likes).Count(&count)
		countData = append(countData, &CountData{v, count})
	}
	return countData
}

// ログイン中のユーザーがいいね済みかチェックする
func CheckFavorite(articleIds []int, userId int) []*Favoritedata {
	db := gormConnect()
	var likes []entity.Likes
	var count int
	favoriteData := []*Favoritedata{}

	for _, v := range articleIds {
		db.Where("article_id = ? AND user_id = ?", v, userId).Find(&likes).Count(&count)
		if count == 1 {
			favoriteData = append(favoriteData, &Favoritedata{v, false})
		} else {
			favoriteData = append(favoriteData, &Favoritedata{v, true})
		}
	}
	return favoriteData
}