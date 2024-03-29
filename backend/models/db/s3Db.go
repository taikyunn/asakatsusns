package db

import (
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func UploadS3Bucket(file *multipart.FileHeader, filename string) {
	sess := Credentials()

	src, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	bucketName := os.Getenv("BUCKET_NAME")
	objectKey := "images/" + filename

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   src,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func DownloadS3Bucket(filepath string) (*os.File, string) {
	sess := Credentials()

	bucket := os.Getenv("BUCKET_NAME")
	key := filepath

	pos := strings.LastIndex(key, ".")
	var filename string
	var extension string

	if (key[pos:]) == ".jpeg" {
		extension = "jpeg"
		filename = "download1.jpeg"
	} else if (key[pos:]) == ".jpg" {
		extension = "jpg"
		filename = "download1.jpg"
	} else {
		extension = "png"
		filename = "download1.png"
	}
	log.Println(filename)

	f, err := os.Create(filename)
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

	return f, extension
}

// 認証情報の関数
func Credentials() *session.Session {
	var AWS_ACCESS_KEY string
	var AWS_SECRET_ACCESS_KEY string

	err := godotenv.Load("env/dev.env")
	if err != nil {
		AWS_ACCESS_KEY = os.Getenv("AWS_ACCESS_KEY")
		AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	} else {
		AWS_ACCESS_KEY = os.Getenv("AWS_ACCESS_KEY")
		AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	}

	creds := credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_ACCESS_KEY, "")

	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-northeast-1")},
	)
	if err != nil {
		log.Fatal(err)
	}

	return sess
}
