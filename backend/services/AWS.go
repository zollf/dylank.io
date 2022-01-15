package services

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GetSession() (*session.Session, error) {
	log.Printf("Getting AWS session")

	aws_region := os.Getenv("AWS_REGION")
	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")

	log.Printf("%s", aws_access_key_id)
	return session.NewSession(&aws.Config{
		Region: aws.String(aws_region),
		Credentials: credentials.NewStaticCredentials(
			aws_access_key_id, aws_secret_access_key, "",
		),
	})
}

func UploadImageToS3(file *multipart.FileHeader, title string) (string, error) {
	bucket := os.Getenv("S3_BUCKET")
	session, err := GetSession()

	log.Printf("Uploading file to s3")

	uploader := s3manager.NewUploader(session)

	body, err := file.Open()
	defer body.Close()

	if err != nil {
		return "", err
	}

	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(title),
		Body:   body,
	})

	if err != nil {
		return "", err
	}

	return upload.Location, nil
}