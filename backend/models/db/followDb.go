package db

import (
	entity "app/models/entity"
)

type FollowList struct {
	UserId           int
	Name             string
	ProfileImagePath string
}

// フォローしているか判別
func CheckFollow(followerId int, followedId int) bool {
	db := gormConnect()
	var follow []entity.Follow
	var count int

	if err := db.Model(&follow).Where("follower_id = ? AND followed_id = ?", followerId, followedId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
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

	if err := db.Where("follower_id = ? AND followed_id = ?", followerId, followedId).Delete(&follow).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

// フォロワー数取得
func GetFollower(userId int) int {
	db := gormConnect()
	var follow []entity.Follow
	var count int

	if err := db.Model(&follow).Where("follower_id = ?", userId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return count
}

// フォロー数取得
func GetFollow(userId int) int {
	db := gormConnect()
	var follow []entity.Follow
	var count int

	if err := db.Model(&follow).Where("followed_id = ?", userId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return count
}

// フォロー一覧取得
func GetFollowerList(followerId int) []int {
	db := gormConnect()
	var follow []entity.Follow

	if err := db.Select("followed_id").Where("follower_id = ?", followerId).Find(&follow).Error; err != nil {
		panic(err.Error())
	}

	followedIds := make([]int, len(follow))
	for i, v := range follow {
		followedIds[i] = v.FollowedId
	}
	return followedIds
}

func GetFollowNameList(userIds []int) []*FollowList {
	db := gormConnect()
	var user []entity.User
	followList := []*FollowList{}

	for _, v := range userIds {
		if err := db.Select("name, profile_image_path").Where("id = ?", v).Find(&user).Error; err != nil {
			panic(err.Error())
		}
		followList = append(followList, &FollowList{v, user[0].Name, user[0].ProfileImagePath})
	}
	defer db.Close()
	return followList
}

// フォロワーユーザーID
func GetFollowedList(followedId int) []int {
	db := gormConnect()
	var follow []entity.Follow

	if err := db.Select("follower_id").Where("followed_id = ?", followedId).Find(&follow).Error; err != nil {
		panic(err.Error())
	}

	followerIds := make([]int, len(follow))
	for i, v := range follow {
		followerIds[i] = v.FollowerId
	}
	defer db.Close()
	return followerIds
}
