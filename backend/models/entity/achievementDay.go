package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AchievementDay struct {
	gorm.Model
	UserId int       `gorm:"type:int; not null" json:"userId"`
	Date   time.Time `gorm:"type:date" json:"date"`
}
