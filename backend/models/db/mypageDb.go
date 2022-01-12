package db

import entity "app/models/entity"

// ユーザーIDからユーザー情報を取得
func GetOneUser(userID int) []entity.User {
	db := gormConnect()
	var user []entity.User

	db.Where("id = ?", userID).Find(&user, userID)
	defer db.Close()
	return user
}

// 寝る時間の登録
func UpdateSleepTime(sleepTime string, userID int) {
	db := gormConnect()
	var user []entity.User

	db.Model(&user).Where("id = ?", userID).Update("sleep_time", sleepTime)
	defer db.Close()
}

func UpdateWakeUpTime(wakeUpTime string, userID int) {
	db := gormConnect()
	var user []entity.User

	db.Model(&user).Where("id = ?", userID).Update("wake_up_time", wakeUpTime)
	defer db.Close()
}

// 寝る時間を取得
func GetSleepTimeById(userID int) []entity.User {
	db := gormConnect()
	var user []entity.User

	db.Select("sleep_time").Where("id = ?", userID).Find(&user)

	defer db.Close()
	return user
}

// 起きる時間を取得
func GetWakeUpTimeById(userID int) []entity.User {
	db := gormConnect()
	var user []entity.User

	db.Select("wake_up_time").Where("id = ?", userID).Find(&user)

	defer db.Close()
	return user
}
