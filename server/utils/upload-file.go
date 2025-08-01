package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadFile faz o upload de um arquivo para o Cloudflare R2 e retorna uma URL assinada para acesso.
func UploadFile(file *multipart.FileHeader) (string, error) {
	// 1. VERIFICAÇÃO DAS VARIÁVEIS DE AMBIENTE
	bucketName := os.Getenv("BUCKET_NAME")
	accountID := os.Getenv("ACCOUNT_ID")
	accessKeyID := os.Getenv("ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	if bucketName == "" || accountID == "" || accessKeyID == "" || accessKeySecret == "" {
		return "", errors.New("uma ou mais variáveis de ambiente (BUCKET_NAME, ACCOUNT_ID, ACCESS_KEY_ID, ACCESS_KEY_SECRET) não foram definidas")
	}

	// 2. CONFIGURAÇÃO DO ENDPOINT DO R2 (MÉTODO MODERNO)
	resolver := aws.EndpointResolverV2Func(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		r2Endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)
		return aws.Endpoint{
			URL: r2Endpoint,
		}, nil
	})

	log.Println("Carregando configuração da AWS...")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, "")),
		config.WithEndpointResolverV2(resolver),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Printf("ERRO AO CARREGAR CONFIG: %v", err)
		return "", fmt.Errorf("erro ao carregar configuração da AWS: %w", err)
	}

	// 3. UPLOAD DO ARQUIVO
	client := s3.NewFromConfig(cfg)

	src, err := file.Open()
	if err != nil {
		log.Printf("ERRO AO ABRIR ARQUIVO: %v", err)
		return "", fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer src.Close()

	log.Printf("Iniciando upload do arquivo '%s' para o bucket '%s'...", file.Filename, bucketName)
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(file.Filename),
		Body:        src,
		ContentType: aws.String(file.Header.Get("Content-Type")),
	})
	if err != nil {
		log.Printf("ERRO NO UPLOAD PARA O R2: %v", err)
		return "", fmt.Errorf("não foi possível fazer o upload do arquivo para o R2: %w", err)
	}
	log.Println("Upload concluído com sucesso.")

	// 4. GERAÇÃO DA URL ASSINADA
	log.Println("Gerando URL assinada...")
	presignClient := s3.NewPresignClient(client)
	presignedReq, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Filename),
	}, func(po *s3.PresignOptions) {
		po.Expires = 15 * time.Minute
	})
	if err != nil {
		log.Printf("ERRO AO GERAR URL ASSINADA: %v", err)
		return "", fmt.Errorf("não foi possível assinar a URL do objeto: %w", err)
	}

	log.Printf("URL gerada: %s", presignedReq.URL)
	return presignedReq.URL, nil
}

