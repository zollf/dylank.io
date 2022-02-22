package services

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetSession() (*session.Session, error) {
	log.Printf("Getting AWS session")

	aws_region := os.Getenv("AWS_REGION")
	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")

	return session.NewSession(&aws.Config{
		Region: aws.String(aws_region),
		Credentials: credentials.NewStaticCredentials(
			aws_access_key_id, aws_secret_access_key, "",
		),
	})
}

func UploadImageToS3(file *multipart.FileHeader, title string) (string, error) {
	bucket := os.Getenv("S3_BUCKET")
	region := os.Getenv("AWS_REGION")
	session, err := GetSession()

	log.Printf("Preparing to upload file")

	s3Client := s3.New(session)

	log.Printf("Opening files content")

	body, err := file.Open()
	defer body.Close()

	buffer := make([]byte, file.Size)
	body.Read(buffer)

	if err != nil {
		return "", err
	}

	log.Printf("Uploading file to s3")

	upload, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(title),
		Body:        bytes.NewReader(buffer),
		ContentType: aws.String(file.Header.Get("Content-Type")),
	})

	log.Printf("Finished Uploading")

	if err != nil {
		return "", err
	}

	log.Printf("%v", upload)

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, title), nil
}
