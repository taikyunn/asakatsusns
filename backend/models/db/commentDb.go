package db

import "app/models/entity"

// コメントを登録
func InsertComment(comment *entity.Comment) {
	db := gormConnect()

	db.Create(&comment)
	defer db.Close()
}
