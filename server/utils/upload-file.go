package utils

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.comcom/aws/aws-sdk-go-v2/service/s3"
)

// UploadFile faz o upload de um arquivo para o Cloudflare R2 e retorna uma URL assinada para acesso.
func UploadFile(file *multipart.FileHeader) (string, error) {
	// Carrega as credenciais das variáveis de ambiente
	bucketName := os.Getenv("BUCKET_NAME")
	accountID := os.Getenv("ACCOUNT_ID")
	accessKeyID := os.Getenv("ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	resolver := aws.EndpointResolverV2Func(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		r2Endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)
		if service == s3.ServiceID {
			return aws.Endpoint{
				URL:               r2Endpoint,
				HostnameImmutable: true,
				Source:            aws.EndpointSourceCustom,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolverV2(resolver), // Usando a função atualizada
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, "")),
		config.WithRegion("auto"), // A região é "auto" para o R2
	)
	if err != nil {
		log.Printf("erro ao carregar configuração da AWS: %v", err)
		return "", fmt.Errorf("erro ao carregar configuração da AWS: %w", err)
	}
	client := s3.NewFromConfig(cfg)

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer src.Close()

	uploaderInput := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(file.Filename),
		Body:        src,
		ContentType: aws.String(file.Header.Get("Content-Type")),
	}

	_, err = client.PutObject(context.TODO(), uploaderInput)
	if err != nil {
		return "", fmt.Errorf("não foi possível fazer o upload do arquivo para o R2: %w", err)
	}

	presignClient := s3.NewPresignClient(client)

	presignedUrl, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Filename),
	}, func(po *s3.PresignOptions) {
		// Define a validade da URL
		po.Expires = 15 * time.Minute
	})
	if err != nil {
		return "", fmt.Errorf("não foi possível assinar a URL do objeto: %w", err)
	}

	return presignedUrl.URL, nil
}

