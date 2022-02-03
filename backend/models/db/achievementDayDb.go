package db

import (
	entity "app/models/entity"
	"time"
)

func RegisterAchievementDay(date time.Time, userId uint) {
	db := gormConnect()
	var achievementDay = entity.AchievementDay{
		UserId: int(userId),
		Date:   date,
	}

	db.Create(&achievementDay)
}
