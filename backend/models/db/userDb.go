package db

import (
	"github.com/jinzhu/gorm"

	entity "app/models/entity"
)

// User モデルの宣言
type User struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}

// ユーザーデータ登録
func InsertUser(signUp *entity.User) {
	db := gormConnect()

	db.Create(&signUp)
	defer db.Close()
}

// ユーザーデータの取得
func GetUserData() []entity.User {
	db := gormConnect()
	var user []entity.User

	if err := db.Select("id,name").Find(&user).Error; err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return user
}

// 正しい名前とパスワードの組み合わせかDBでチェック
func CheckNameAndPassword(name string, email string) bool {
	db := gormConnect()
	var user User

	if err := db.Where("name = ? AND email = ?", name, email).Find(&user).Error; err != nil {
		return false
	}
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

	if err := db.Select("name").Where("id = ?", userID).Find(&user).Error; err != nil {
		panic(err.Error())
	}
	return user
}
