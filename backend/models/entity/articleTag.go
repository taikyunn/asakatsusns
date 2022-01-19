package entity

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	gorm.Model
	ArticleId int `gorm:"type:int; not null constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"articleId"`
	TagId     int `gorm:"type:int; not null constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tagId"`
}
