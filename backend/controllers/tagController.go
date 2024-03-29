package controller

import (
	db "app/models/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AutoCompleteItems struct {
	Text string `json:"text"`
}

type TagResult struct {
	ArticleId int
	Key       uint
	Value     string
}

type ArticleResult struct {
	Id        int
	Name      string
	Body      string
	UserId    int
	UpdatedAt string
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

// タグごとの記事データを取得
func GetTagArticles(c *gin.Context) {
	tagIdStr := c.PostForm("tagId")
	tagId, _ := strconv.Atoi(tagIdStr)

	// tagIdをもとにarticleIdを取得
	articleIds := db.GetArticleIdByTagId(tagId)

	// articleIdsを元に記事データを取得
	articles := db.GetArticleByTag(articleIds)

	// タグ情報の取得
	tagInfo := db.GetTagInfo(articleIds)

	tagMap := make(map[uint]string, len(tagInfo))
	for _, v := range tagInfo {
		for i := 0; i < len(v.Name); i++ {
			tagMap[v.TagId[i]] = v.Name[i]
		}
	}

	tagResult := []*TagResult{}

	for _, v := range tagInfo {
		for i := 0; i < len(v.TagId); i++ {
			for key := range tagMap {
				if (v.TagId[i]) == key {
					tagResult = append(tagResult, &TagResult{v.ArticleId, key, tagMap[key]})
				}
			}
		}
	}

	// Idに紐づいた中身を取得
	tag := db.GetTagNameById(tagId)

	// Idに紐づいたタグの件数を取得
	count := db.GetTagCount(tagId)

	c.JSON(200, gin.H{"article": articles, "tag": tagResult, "topTag": tag, "count": count})
}
