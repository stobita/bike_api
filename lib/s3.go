package lib

import (
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Upload(data io.ReadSeeker, fileName string) string {
	client := s3.New(session.New(), &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			"sample_access_key",
			"sample_secret_key",
			"",
		),
		Region:           aws.String("sample_region"),
		Endpoint:         aws.String("http://minio:9000"),
		S3ForcePathStyle: aws.Bool(true),
	})
	_, err := client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("bike-api-sample"),
		Key:    aws.String(fileName),
		Body:   data,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return fmt.Sprintf("http://localhost:9000/bike-api-sample/" + fileName)
}
