package db

import (
	"app/models/entity"
	"time"
)

type CommentData struct {
	Name             string
	Comment          string
	ProfileImagePath string
	UpdatedAt        time.Time
}

type ResultCommentData struct {
	Name             string
	Comment          string
	ProfileImagePath string
	Image            string
	UpdatedAt        string
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
func GetCommentData(articleId int) []*ResultCommentData {
	db := gormConnect()
	commentData := []*CommentData{}
	resultCommentData := []*ResultCommentData{}
	cols := "name, comment, profile_image_path, comment.updated_at"
	table := "INNER JOIN user ON comment.user_id = user.id"
	where := "comment.deleted_at IS NULL AND user.deleted_at IS NULL AND article_id = ?"

	if err := db.Table("comment").Select(cols).Joins(table).Where(where, articleId).Find(&commentData).Error; err != nil {
		panic(err.Error())
	}
	for _, v := range commentData {
		t := v.UpdatedAt.Format("2006/01/02 15:04:05")
		resultCommentData = append(resultCommentData, &ResultCommentData{v.Name, v.Comment, v.ProfileImagePath, "", t})
	}
	defer db.Close()
	return resultCommentData
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
	defer db.Close()
	return commentCount
}

// コメント件数を取得(1件)
func GetOneCommentCount(articleId int) int {
	db := gormConnect()
	var comment []entity.Comment
	var count int

	if err := db.Where("article_id = ?", articleId).Find(&comment).Count(&count).Error; err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return count
}
