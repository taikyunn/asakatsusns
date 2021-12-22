package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name             string  `form:"name" binding:"required" gorm:"type:varchar(200); not null" json:"name"`
	ProfileImage     *string `gorm:"type:varchar(200) DEFAULT NULL" json:"ProfileImage"`
	SelfIntroduction *string `gorm:"type:varchar(200) DEFAULT NULL" json:"SelfIntroduction"`
	WakeUpTime       *string `gorm:"type:varchar(200) DEFAULT NULL" json:"WakeUpTime"`
	SleepTime        *string `gorm:"type:varchar(200) DEFAULT NULL" json:"SleepTime"`
	RangeOfSuccess   int     `gorm:"type:int" json:"RangeOfSuccess"`
	Email            string  `gorm:"type:varchar(200);not null" json:"email"`
	EmailVerifiedAt  bool    `gorm:"type:bool; DEFAULT 0" json:"EmailVerifiedAt"`
	Password         string  `form:"password" binding:"required" gorm:"type:varchar(200)" json:"password"`
}

type Password struct {
	Password string `gorm:"type:varchar(200)" json:"password"`
}
