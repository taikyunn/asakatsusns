package controller

import (
	db "app/models/db"
	"app/models/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

// いいね登録
func RegisterLikes(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)

	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	var likes = entity.Likes{
		ArticleId: articleID,
		UserId:    userID,
	}

	db.RegisterLikes(&likes)
}

// いいね解除
func DeleteLikes(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)

	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	db.DeleteLikes(articleID, userID)
}

// いいね数の取得
func GetCountFavorites(c *gin.Context) {
	// 直近10件の投稿データのarticleIdを取得
	articleIds := db.GetLastTenArticleID()

	articleID := make([]int, len(articleIds))
	for i, v := range articleIds {
		articleID[i] = int(v.ID)
	}

	// いいね件数を取得
	countData := db.GetLikeCount(articleID)
	c.JSON(200, countData)
}

// 無限スクロール用・いいね数取得
func GetNextCountFavorites(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)

	// いいね数を取得
	countData := db.GetPreviousLikeCount(articleID)
	c.JSON(200, countData)
}

func GetOneCountFavorites(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)

	count := db.GetOneLikeCount(articleID)
	c.JSON(200, count)
}

// ログイン中のユーザーのいいね状態の取得
func CheckFavorite(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	articleIds := db.GetLastTenArticleID()
	articleID := make([]int, len(articleIds))
	for i, v := range articleIds {
		articleID[i] = int(v.ID)
	}

	favoriteData := db.CheckFavorite(articleID, userID)
	c.JSON(200, favoriteData)
}

// 無限スクロール・いいね状態の取得
func CheckNextFavorite(c *gin.Context) {
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)

	favoriteData := db.CheckNextFavorite(articleID, userID)
	c.JSON(200, favoriteData)
}

// 詳細ページ・ログイン中のユーザーのいいね状態の取得
func CheckFavoriteByArticleId(c *gin.Context) {
	articleIdStr := c.PostForm("articleId")
	articleID, _ := strconv.Atoi(articleIdStr)
	userIdStr := c.PostForm("userId")
	userID, _ := strconv.Atoi(userIdStr)

	favoriteData := db.CheckFavoriteByArticleId(articleID, userID)
	c.JSON(200, favoriteData)
}

// いいねした記事最新10件を取得
func GetLikedPost(c *gin.Context) {
	mypageUserIdStr := c.PostForm("mypageUserId")
	mypageUserId, _ := strconv.Atoi(mypageUserIdStr)

	// いいねした記事のidを取得
	articleIDs := db.GetLikedPostId(mypageUserId)

	// いいね記事の中身を取得
	favoritePostData := db.GetLikedPost(articleIDs)

	// コメント件数取得
	commentCount := db.GetCommentCount(articleIDs)

	// タグ情報の取得
	tagInfo := db.GetTagInfo(articleIDs)

	tagMap := make(map[uint]string, len(tagInfo))
	for _, v := range tagInfo {
		for i := 0; i < len(v.Name); i++ {
			tagMap[v.TagId[i]] = v.Name[i]
		}
	}

	dbMypageTagResult := []*DbMypageTagResult{}

	for _, v := range tagInfo {
		for i := 0; i < len(v.TagId); i++ {
			for key := range tagMap {
				if (v.TagId[i]) == key {
					dbMypageTagResult = append(dbMypageTagResult, &DbMypageTagResult{v.ArticleId, key, tagMap[key]})
				}
			}
		}
	}

	c.JSON(200, gin.H{"favoritePostData": favoritePostData, "commentCount": commentCount, "tagData": dbMypageTagResult})
}

// いいねしているか判定(いいねした投稿)
func CheckFavoriteLikedPost(c *gin.Context) {
	mypageUserIdStr := c.PostForm("mypageUserId")
	mypageUserId, _ := strconv.Atoi(mypageUserIdStr)
	visiterUserIdStr := c.PostForm("visiterUserId")
	visiterUserID, _ := strconv.Atoi(visiterUserIdStr)

	// いいねした記事のidを取得
	articleIDs := db.GetLikedPostId(mypageUserId)

	// いいね済みか判別
	favoriteData := db.CheckFavorite(articleIDs, visiterUserID)

	c.JSON(200, favoriteData)
}

func GetCountFavoriteLikedPost(c *gin.Context) {
	mypageUserIdStr := c.PostForm("mypageUserId")
	mypageUserId, _ := strconv.Atoi(mypageUserIdStr)

	// いいねした記事のidを取得
	articleIDs := db.GetLikedPostId(mypageUserId)

	// いいね記事の中身を取得
	countData := db.GetLikedPostCount(articleIDs)

	c.JSON(200, countData)
}
