package controller

import (
	db "app/models/db"

	"github.com/gin-gonic/gin"
)

type AutoCompleteItems struct {
	Text string `json:"text"`
}

// 自動補完データの取得
func GetAutocompleteItems(c *gin.Context) {
	tagData := db.GetAutocompleteItems()

	tagNames := make([]string, len(tagData))
	for i, v := range tagData {
		tagNames[i] = v.Name
	}

	c.JSON(200, tagNames)
}
