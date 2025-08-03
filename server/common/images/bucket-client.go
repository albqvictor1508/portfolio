package images

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	awsRegion = "auto"
)

func NewBucket() (*s3.Client, error) {
	accountID := os.Getenv("R2_ACCOUNT_ID")
	accessKeyID := os.Getenv("R2_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("R2_ACCESS_KEY_SECRET")

	if accountID == "" || accessKeyID == "" || accessKeySecret == "" {
		return nil, fmt.Errorf("R2 environment variables not set (R2_ACCOUNT_ID, R2_ACCESS_KEY_ID, R2_ACCESS_KEY_SECRET)")
	}

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, "")),
		config.WithRegion(awsRegion),
	)
	if err != nil {
		log.Printf("failed to load SDK config: %v", err)
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}

func UploadFile(fileHeader *multipart.FileHeader, path string) (string, error) {
	client, err := NewBucket()
	if err != nil {
		return "", fmt.Errorf("failed to create S3 client: %w", err)
	}

	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		return "", fmt.Errorf("BUCKET_NAME environment variable not set")
	}

	r2PublicURL := os.Getenv("R2_PUBLIC_URL")
	if r2PublicURL == "" {
		return "", fmt.Errorf("R2_PUBLIC_URL environment variable not set")
	}

	f, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
		Body:   f,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to R2: %w", err)
	}

	url := fmt.Sprintf("https://%s/%s", r2PublicURL, path)

	return url, nil
}
