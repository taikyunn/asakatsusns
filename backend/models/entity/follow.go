package entity

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	FollowerId int `gorm:"type:int not null" json:"followerId"`
	FollowedId int `gorm:"type:int not null" json:"followedId"`
}
