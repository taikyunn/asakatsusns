package entity

import (
	"github.com/jinzhu/gorm"
)

type AchievementDay struct {
	gorm.Model
	UserId int `gorm:"type:int; not null" json:"userId"`
}
