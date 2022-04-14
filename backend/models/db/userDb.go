package db

import (
	entity "app/models/entity"
)

// 名前の重複チェック
func CheckNameUnique(name string) (ok bool, result map[string]string) {
	db := gormConnect()
	var user []entity.User
	var count int
	result = make(map[string]string)

	if err := db.Where("name = ?", name).Find(&user).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	if count == 1 {
		result["Name"] = "*この名前は既に登録されています。"
		return false, result
	}
	return true, result
}

// ユーザーデータ登録
func InsertUser(signUp *entity.User) {
	db := gormConnect()

	db.Create(&signUp)
	defer db.Close()
}

// 正しい名前とパスワードの組み合わせかDBでチェック
func CheckNameAndPassword(name string, email string) bool {
	db := gormConnect()
	var user []entity.User

	if err := db.Where("name = ? AND email = ?", name, email).Find(&user).Error; err != nil {
		return false
	}
	defer db.Close()
	return true
}

// ユーザーIDの配列からユーザー名を取得
func GetNameById(userID []uint) []entity.User {
	db := gormConnect()
	var user []entity.User

	db.Select("id,name").Find(&user, userID)
	defer db.Close()
	return user
}

// ユーザーIdからユーザー名を取得(単数)
func GetUserName(userID uint) []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Select("name, profile_image_path").Where("id = ?", userID).Find(&user).Error; err != nil {
		panic(err.Error())
	}
	return user
}

func GetImagePath(userId int) []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Select("profile_image_path").Where("id = ?", userId).Find(&user).Error; err != nil {
		panic(err.Error())
	}
	return user
}
