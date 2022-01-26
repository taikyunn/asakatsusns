package db

import (
	entity "app/models/entity"
)

func CheckFollow(followerId int) bool {
	db := gormConnect()
	var follow []entity.Follow
	var count int

	if err := db.Model(&follow).Where("followed_id = ?", followerId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	if count == 0 {
		return false
	} else {
		return true
	}
}
