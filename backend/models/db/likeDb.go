package db

import entity "app/models/entity"

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
