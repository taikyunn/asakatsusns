package db

import (
	"app/models/entity"
	"strings"
	"time"
)

type Ranking struct {
	Name  string
	Count int
}

func RegisterAchievementDay(date time.Time, userId uint) int {
	const layout = "2006-01-02"
	createdAtStr := date.Format(layout)
	createdAt := strings.Split(createdAtStr, " ")
	time := createdAt[0]

	db := gormConnect()
	var count int
	var achievementDay []entity.AchievementDay

	if err := db.Model(&achievementDay).Where("created_at LIKE ?", time+" %%:%%:%%").Count(&count).Error; err != nil {
		panic(err.Error())
	}
	if count == 1 {
		return 0
	}
	var insertData = entity.AchievementDay{
		UserId: int(userId),
	}
	db.Create(&insertData)

	if err := db.Model(&achievementDay).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		panic(err.Error())
	}
	return count
}

// ランキング取得
func GetWakeUpRanking() []*Ranking {
	db := gormConnect()
	var count int
	var user []entity.User
	ranking := []*Ranking{}

	if err := db.Select("name").Joins("INNER JOIN achievement_day ON achievement_day.user_id = user.id").Group("achievement_day.user_id").Order("COUNT(user_id) DESC").Limit(5).Find(&user).Error; err != nil {
		panic(err.Error())
	}

	userName := make([]string, len(user))
	for i, v := range user {
		userName[i] = v.Name
	}

	for _, v := range userName {
		if err := db.Table("achievement_day").Select("count(achievement_day.id)").Joins("INNER JOIN user ON achievement_day.user_id = user.id").Where("name = ?", v).Count(&count).Error; err != nil {
			panic(err.Error())
		}
		ranking = append(ranking, &Ranking{v, count})
	}
	return ranking
}
