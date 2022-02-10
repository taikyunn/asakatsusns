package db

import (
	"app/models/entity"
)

type CommentData struct {
	Name    string
	Comment string
}

type CommentCount struct {
	ArticleId int
	Count     int
}

// コメントを登録
func InsertComment(comment *entity.Comment) {
	db := gormConnect()

	db.Create(&comment)
	defer db.Close()
}

// コメント情報取得
func GetCommentData(articleId int) []*CommentData {
	db := gormConnect()
	var comment []entity.Comment
	var user []entity.User

	if err := db.Select("user_id, comment").Where("article_id = ?", articleId).Find(&comment).Error; err != nil {
		panic(err.Error())
	}

	comments := make([]string, len(comment))
	userIDs := make([]int, len(comment))
	for i, v := range comment {
		comments[i] = v.Comment
		userIDs[i] = v.UserId
	}

	// 誰がコメントしたかを取得
	if err := db.Select("id,name").Where("id IN (?)", userIDs).Find(&user).Error; err != nil {
		panic(err.Error())
	}

	commentData := []*CommentData{}

	for _, cv := range comment {
		for _, uv := range user {
			if cv.UserId == int(uv.ID) {
				commentData = append(commentData, &CommentData{uv.Name, cv.Comment})
			}
		}
	}
	return commentData
}

// コメント件数を取得
func GetCommentCount(articleIds []int) []*CommentCount {
	db := gormConnect()
	var comment []entity.Comment
	var count int
	commentCount := []*CommentCount{}

	for _, v := range articleIds {
		if err := db.Where("article_id = ?", v).Find(&comment).Count(&count).Error; err != nil {
			panic(err.Error())
		}
		commentCount = append(commentCount, &CommentCount{v, count})
	}
	return commentCount
}
