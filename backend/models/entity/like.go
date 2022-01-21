package entity

import "github.com/jinzhu/gorm"

type Likes struct {
	gorm.Model
	UserId    int `gorm:"type:int not null" json:"userId"`
	ArticleId int `gorm:"type:int not null" json:"articleId"`
}
