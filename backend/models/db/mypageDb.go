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

// 画像ファイルのパスを保存
func UpdateFilePath(id int, file_path string) {
	db := gormConnect()
	var user []entity.User

	db.Model(&user).Where("id = ?", id).Update("profile_image_path", file_path)
	defer db.Close()
}

// profile_image_pathのデータの取得
func GetFilePathById(id int) []entity.User {
	db := gormConnect()
	var user []entity.User

	db.Select("profile_image_path").Where("id = ?", id).Find(&user)
	defer db.Close()
	return user
}

// ユーザー名の更新
func UpdateUserName(id int, name string) {
	db := gormConnect()
	var user []entity.User

	db.Model(&user).Where("id = ?", id).Update("name", name)
	defer db.Close()
}

// 投稿取得(マイページ)
func GetMypageArticle(userID int) ([]entity.Article, []int) {
	db := gormConnect()
	var article []entity.Article

	db.Limit(10).Order("id DESC").Select("id, body").Where("user_id = ?", userID).Find(&article)
	defer db.Close()

	articleIds := make([]int, len(article))
	for i, v := range article {
		articleIds[i] = int(v.ID)
	}
	return article, articleIds
}
