package db

import (
	entity "app/models/entity"
)

type WakeUpData struct {
	WakeUpTime     string
	RangeOfSuccess int
}

// ユーザーIDからユーザー情報を取得
func GetOneUser(userID int) []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Where("id = ?", userID).Find(&user, userID).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return user
}

// 寝る時間の登録
func UpdateSleepTime(sleepTime string, userID int) {
	db := gormConnect()
	var user []entity.User

	if err := db.Model(&user).Where("id = ?", userID).Update("sleep_time", sleepTime).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func UpdateWakeUpTime(wakeUpTime string, userID int) {
	db := gormConnect()
	var user []entity.User

	if err := db.Model(&user).Where("id = ?", userID).Update("wake_up_time", wakeUpTime).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

// 寝る時間を取得
func GetSleepTimeById(userID int) []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Select("sleep_time").Where("id = ?", userID).Find(&user).Error; err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return user
}

// 起きる時間を取得
func GetWakeUpTimeById(userID int) []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Select("wake_up_time").Where("id = ?", userID).Find(&user).Error; err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return user
}

// 起きる時間を取得
func GetWakeUpData(userID int) []*WakeUpData {
	db := gormConnect()
	wakeUpData := []*WakeUpData{}

	if err := db.Table("user").Select("wake_up_time, range_of_success").Where("id = ?", userID).Scan(&wakeUpData).Error; err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return wakeUpData
}

// 画像ファイルのパスを保存
func UpdateFilePath(id int, file_path string) {
	db := gormConnect()
	var user []entity.User

	if err := db.Model(&user).Where("id = ?", id).Update("profile_image_path", file_path).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

// profile_image_pathのデータの取得
func GetFilePathById(id int) []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Select("profile_image_path").Where("id = ?", id).Find(&user).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return user
}

// ユーザー名の更新
func UpdateUserName(id int, name string) {
	db := gormConnect()
	var user []entity.User

	if err := db.Model(&user).Where("id = ?", id).Update("name", name).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

// 投稿取得(マイページ)
func GetMypageArticle(userID int) ([]entity.Article, []int) {
	db := gormConnect()
	var article []entity.Article

	if err := db.Limit(10).Order("id DESC").Select("id, body").Where("user_id = ?", userID).Find(&article).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()

	articleIds := make([]int, len(article))
	for i, v := range article {
		articleIds[i] = int(v.ID)
	}
	return article, articleIds
}
