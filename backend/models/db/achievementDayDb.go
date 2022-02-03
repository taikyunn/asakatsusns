package db

import (
	"app/models/entity"
	"strings"
	"time"
)

func RegisterAchievementDay(date time.Time, userId uint) {
	const layout = "2006-01-02"
	createdAtStr := date.Format(layout)
	createdAt := strings.Split(createdAtStr, " ")
	time := createdAt[0]

	db := gormConnect()
	var count int64
	var achievementDay []entity.AchievementDay

	if err := db.Model(&achievementDay).Where("created_at LIKE ?", time+" %%:%%:%%").Count(&count).Error; err != nil {
		panic(err.Error())
	}
	if count == 1 {
		return
	}
	var insertData = entity.AchievementDay{
		UserId: int(userId),
	}

	db.Create(&insertData)
}
