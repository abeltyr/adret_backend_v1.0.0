package s3Config

import (
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Upload(
	key string,
	body io.ReadSeeker,
	acl string,
) (*s3.PutObjectOutput, error) {

	bucket := os.Getenv("AWS_BUCKET")

	s3Client, err := GetS3Client()
	if err != nil {
		return nil, err
	}

	object := s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
		// Metadata: map[string]*string{
		// 	"x-amz-meta-my-key": aws.String(key),
		// },
	}

	if acl != "" {
		object.ACL = aws.String(acl)
	}

	data, err := s3Client.PutObject(&object)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	return data, nil
}
