package s3Config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func MultipleDelete(keys []*s3.ObjectIdentifier) (*s3.DeleteObjectsOutput, error) {

	bucket := os.Getenv("AWS_BUCKET")

	s3Client, err := GetS3Client()
	if err != nil {
		return nil, err
	}

	object := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucket),
		Delete: &s3.Delete{
			Objects: keys,
		},
	}
	log.Println(object)

	data, err := s3Client.DeleteObjects(object)

	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	return data, nil
}
