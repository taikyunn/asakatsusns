package entity

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	gorm.Model
	ArticleId int `gorm:"type:int; not null" json:"articleId"`
	TagId     int `gorm:"type:int; not null" json:"tagId"`
}
