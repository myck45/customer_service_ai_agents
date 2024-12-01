package providers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewAWSS3Client(region string) *s3.S3 {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	return s3.New(sess)
}
