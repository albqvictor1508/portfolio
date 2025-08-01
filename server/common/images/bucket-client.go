package images

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewBucketClient() *s3.Client {
	bucketName := os.Getenv("BUCKET_NAME")
	accountId := os.Getenv("ACCOUNT_ID")
	accessKeyId := os.Getenv("ACCESS_KEY_ID")
	endpoint := os.Getenv("R2_ENDPOINT")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")
	region := os.Getenv("R2_REGION")

	cfg := aws.Config{
		Region:      region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:               endpoint,
					HostnameImmutable: true,
				}, nil
			},
		),
	}
	return s3.NewFromConfig(cfg)
}
