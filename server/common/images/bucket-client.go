package images

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Session *s3.S3

func init() {
	s3Session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})))
}

func ListBuckets() (resp *s3.ListBucketsOutput) {
}
