package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	name := c.PostForm("name")
	log.Println("bodyの中身", name)
	// user_idの取得

	// データの登録
}
