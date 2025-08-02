package images

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewBucketClient() *s3.Client {
	accessKeyId := os.Getenv("ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")
	endpoint := os.Getenv("R2_ENDPOINT")
	region := os.Getenv("R2_REGION")

	ctx := context.Background()
}
