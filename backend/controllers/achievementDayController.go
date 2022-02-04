package controller

import (
	db "app/models/db"

	"github.com/gin-gonic/gin"
)

// ランキング取得
func GetWakeUpRanking(c *gin.Context) {
	ranking := db.GetWakeUpRanking()
	c.JSON(200, ranking)
}
