package entity

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserId    int    `gorm:"type:int; not null" json:"userId"`
	ArticleId int    `gorm:"type:int; not null" json:"articleId"`
	Comment   string `gorm:"type:varchar(200);not null" json:"comment"`
}
