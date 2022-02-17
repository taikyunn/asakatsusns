package middleware

import (
	"app/models/db"
	"context"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bucket string
		var key string

		err := godotenv.Load("env/dev.env")
		if err != nil {
			bucket = os.Getenv("BUCKET_NAME")
			key = os.Getenv("FILE_PATH")
		} else {
			bucket = os.Getenv("BUCKET_NAME")
			key = os.Getenv("FILE_PATH")
		}
		// sessionの作成
		sess := db.Credentials()

		// S3オブジェクトを書き込むファイルの作成
		f, err := os.Create("firebase-admin-sdk.json")
		if err != nil {
			log.Fatal(err)
		}

		// Downloaderを作成し、S3オブジェクトをダウンロード
		downloader := s3manager.NewDownloader(sess)
		n, err := downloader.Download(f, &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("DownloadedSize: %d byte", n)

		// Firebase SDK のセットアップ
		opt := option.WithCredentialsFile("firebase-admin-sdk.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Printf("error: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Printf("error: %v\n", err)
			os.Exit(1)
		}

		// クライアントから送られてきた JWT 取得
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(201, gin.H{"message": "エラー"})
			c.Abort()
		} else {
			log.Printf("Verified ID token: %v\n", token)
			c.Next()
		}
	}
}
