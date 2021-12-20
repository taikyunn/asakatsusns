package entity

import (
	"time"
)

type User struct {
	ID               int       `gorm:"primary_key;not null"       json:"id"`
	Name             string    `gorm:"type:varchar(200); not null" json:"name"`
	ProfileImage     *string   `gorm:"type:varchar(200) DEFAULT NULL" json:"ProfileImage"`
	SelfIntroduction *string   `gorm:"type:varchar(200) DEFAULT NULL" json:"SelfIntroduction"`
	WakeUpTime       *string   `gorm:"type:varchar(200) DEFAULT NULL" json:"WakeUpTime"`
	SleepTime        *string   `gorm:"type:varchar(200) DEFAULT NULL" json:"SleepTime"`
	RangeOfSuccess   int       `gorm:"type:int" json:"RangeOfSuccess"`
	Email            string    `gorm:"type:varchar(200);not null" json:"email"`
	EmailVerifiedAt  bool      `gorm:"type:bool; DEFAULT 0" json:"EmailVerifiedAt"`
	Password         string    `gorm:"type:varchar(200)" json:"password"`
	CreatedAt        time.Time `json:"createdAt" sql:"DEFAULT:current_timestamp"`
	UpdatedAt        time.Time `json:"updatedAt" sql:"DEFAULT:current_timestamp"`
}
