package s3Config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Client() (*s3.S3, error) {

	key := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	location := os.Getenv("AWS_REGION")

	session, _ := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Region:           aws.String(location),
		S3ForcePathStyle: aws.Bool(false), //
	})

	s3Client := s3.New(session)

	return s3Client, nil

}
