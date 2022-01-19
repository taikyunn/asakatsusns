package entity

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	UserId uint   `form:"userId" binding:"required" gorm:"type:int; not null" json:"userId"`
	Body   string `gorm:"type:varchar(200);not null" json:"body"`
}

type UserData struct {
	Name string `form:"name"`
	Body string `form:"body"`
}
