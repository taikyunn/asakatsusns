package db

import (
	entity "app/models/entity"
	"time"
)

type CountData struct {
	ArticleId int
	Count     int
}

type Favoritedata struct {
	ArticleId int
	Count     bool
}

type FavoriteCountData struct {
	ArticleId int
	Count     int
}

type FavoritePostData struct {
	Id               int
	UserId           int
	Name             string
	Body             string
	ProfileImagePath string
	UpdatedAt        time.Time
}

type ResultFavoritePostData struct {
	Id               int
	UserId           int
	Name             string
	Body             string
	ProfileImagePath string
	UpdatedAt        string
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

	if err := db.Where("user_id = ? AND article_id = ? ", userId, articleId).Delete(&likes).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

// 直近10件のarticleIdを取得
func GetLastTenArticleID() []entity.Article {
	db := gormConnect()
	var article []entity.Article

	if err := db.Select("id").Limit(10).Order("id DESC").Find(&article).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return article
}

// いいね数を取得する
func GetLikeCount(articleIds []int) []*CountData {
	db := gormConnect()
	var likes []entity.Likes
	var count int
	countData := []*CountData{}

	for _, v := range articleIds {
		if err := db.Where("article_id = ?", v).Find(&likes).Count(&count).Error; err != nil {
			panic(err.Error())
		}
		countData = append(countData, &CountData{v, count})
	}
	defer db.Close()
	return countData
}

func GetPreviousLikeCount(articleId int) []*CountData {
	db := gormConnect()
	var likes []entity.Likes
	var count int
	countData := []*CountData{}

	if err := db.Where("article_id = ?", articleId).Find(&likes).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	countData = append(countData, &CountData{articleId, count})
	defer db.Close()
	return countData
}

// 次の10件分のarticleIdを取得
func GetNextArticleID(updatedAt time.Time) []int {
	db := gormConnect()
	var article []entity.Article

	if err := db.Select("id").Where("updated_at <  ?", updatedAt).Limit(10).Order("updated_at DESC").Find(&article).Error; err != nil {
		panic(err.Error())
	}
	articleID := make([]int, len(article))
	for i, v := range article {
		articleID[i] = int(v.ID)
	}
	defer db.Close()
	return articleID
}

// ユーザーがいいね済みかチェック
func CheckFavorite(articleIds []int, userId int) []*Favoritedata {
	db := gormConnect()
	var likes []entity.Likes
	var count int
	favoriteData := []*Favoritedata{}

	for _, v := range articleIds {
		if err := db.Where("article_id = ? AND user_id = ?", v, userId).Find(&likes).Count(&count).Error; err != nil {
			panic(err.Error())
		}
		if count == 1 {
			favoriteData = append(favoriteData, &Favoritedata{v, false})
		} else {
			favoriteData = append(favoriteData, &Favoritedata{v, true})
		}
	}
	defer db.Close()
	return favoriteData
}

// ユーザーがいいね済みかチェック(無限スクロール)
func CheckNextFavorite(articleId int, userId int) []*Favoritedata {
	db := gormConnect()
	var likes []entity.Likes
	var count int
	favoriteData := []*Favoritedata{}

	if err := db.Where("article_id = ? AND user_id = ?", articleId, userId).Find(&likes).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	if count == 1 {
		favoriteData = append(favoriteData, &Favoritedata{articleId, false})
	} else {
		favoriteData = append(favoriteData, &Favoritedata{articleId, true})
	}
	defer db.Close()
	return favoriteData
}

// 一件の投稿に対するいいね数を取得
func GetOneLikeCount(articleId int) int {
	db := gormConnect()
	var likes []entity.Likes
	var count int

	if err := db.Where("article_id = ?", articleId).Find(&likes).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return count
}

// ログイン中のユーザーがいいね済みかチェック(単数)
func CheckFavoriteByArticleId(articleId int, userId int) []*Favoritedata {
	db := gormConnect()
	var likes []entity.Likes
	var count int
	favoriteData := []*Favoritedata{}

	if err := db.Where("article_id = ? AND user_id = ?", articleId, userId).Find(&likes).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	if count == 1 {
		favoriteData = append(favoriteData, &Favoritedata{articleId, false})
	} else {
		favoriteData = append(favoriteData, &Favoritedata{articleId, true})
	}
	defer db.Close()
	return favoriteData
}

// 直近10件のいいねした記事IDを取得
func GetLikedPostId(userId int) []int {
	db := gormConnect()
	var likes []entity.Likes

	if err := db.Select("article_id").Limit(10).Order("id DESC").Where("user_id = ?", userId).Find(&likes).Error; err != nil {
		panic(err.Error())
	}

	articleIDs := make([]int, len(likes))
	for i, v := range likes {
		articleIDs[i] = v.ArticleId
	}
	defer db.Close()
	return articleIDs
}

// いいね記事の中身を取得
func GetLikedPost(articleIds []int) []*ResultFavoritePostData {
	db := gormConnect()
	favoritePostData := []*FavoritePostData{}
	resultfavoritePostData := []*ResultFavoritePostData{}

	for _, v := range articleIds {
		if err := db.Table("article").Select("article.id, user_id, name, body, profile_image_path, user.updated_at").Joins("INNER JOIN user ON article.user_id = user.id").Where("article.deleted_at IS NULL AND user.deleted_at IS NULL AND article.id = ?", v).Scan(&favoritePostData).Error; err != nil {
			panic(err.Error())
		}
		for _, v := range favoritePostData {
			t := v.UpdatedAt.Format("2006/01/02 15:04:05")
			resultfavoritePostData = append(resultfavoritePostData, &ResultFavoritePostData{v.Id, v.UserId, v.Name, v.Body, v.ProfileImagePath, t})
		}
	}
	defer db.Close()
	return resultfavoritePostData
}

// いいね記事の中身を取得
func GetLikedPostCount(articleIds []int) []*FavoriteCountData {
	db := gormConnect()
	var likes []entity.Likes
	var count int
	favoriteCountData := []*FavoriteCountData{}

	for _, v := range articleIds {
		if err := db.Where("article_id = ?", v).Find(&likes).Count(&count).Error; err != nil {
			panic(err.Error())
		}
		favoriteCountData = append(favoriteCountData, &FavoriteCountData{v, count})
	}
	defer db.Close()
	return favoriteCountData
}
