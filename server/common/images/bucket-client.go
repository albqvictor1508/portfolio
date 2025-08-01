package images

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type CustomResolver struct{}

func NewBucketClient() *s3.Client {
	accessKeyId := os.Getenv("ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")
	endpoint := os.Getenv("R2_ENDPOINT")
	region := os.Getenv("R2_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithEndpointResolverWithOptions(CustomResolver{Endpoint: endpoint}),
	)
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}

	return s3.NewFromConfig(cfg)
}
