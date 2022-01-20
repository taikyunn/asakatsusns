package entity

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `form:"name" binding:"required" gorm:"type:varchar(200); not null" json:"name"`
}
