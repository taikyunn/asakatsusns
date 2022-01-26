package db

import (
	entity "app/models/entity"
)

// フォローしているか判別
func CheckFollow(followerId int, followedId int) bool {
	db := gormConnect()
	var follow []entity.Follow
	var count int

	if err := db.Model(&follow).Where("follower_id = ? AND followed_id = ?", followerId, followedId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	if count == 0 {
		return false
	} else {
		return true
	}
}

// フォロー登録
func RegisterFollow(followerId int, followedId int) {
	db := gormConnect()
	var follow = entity.Follow{
		FollowerId: followerId,
		FollowedId: followedId,
	}

	db.Create(&follow)
	defer db.Close()
}

// フォロー解除
func DeleteFollow(followerId int, followedId int) {
	db := gormConnect()
	var follow []entity.Follow

	db.Where("follower_id = ? AND followed_id = ?", followerId, followedId).Delete(&follow)
	defer db.Close()
}

// フォロワー数取得
func GetFollower(userId int) int {
	db := gormConnect()
	var follow []entity.Follow
	var count int

	if err := db.Model(&follow).Where("followed_id = ?", userId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	return count
}

// フォロー数取得
func GetFollow(userId int) int {
	db := gormConnect()
	var follow []entity.Follow
	var count int

	if err := db.Model(&follow).Where("follower_id = ?", userId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	return count
}
