package utils

import (
	"fmt"
	"net"
	"time"

	"github.com/joho/godotenv"
	// gomail "gopkg.in/mail.v2"
)

type SendEmailParams struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func SendEmail(params SendEmailParams) error {
	godotenv.Load()
	/*
		var (
			myEmail    = os.Getenv("MY_EMAIL")
			myPassword = os.Getenv("MY_PASSWORD")
		)
	*/
	// message := gomail.NewMessage()

	if params.Subject == "" {
		params.Subject = "By Portfolio"
	}

	conn, err := net.DialTimeout("tcp", "smtp.gmail.com:465", 5*time.Second)
	if err != nil {
		fmt.Println("Erro:", err)
		return fmt.Errorf("esse erro é de subir a connection", err)
	}
	defer conn.Close()

	fmt.Println("Conexão aberta com sucesso!")

	/*
		message.SetHeader("From", params.From)
		message.SetHeader("To", myEmail)
		message.SetHeader("Subject", params.Subject)
		message.SetBody("text/plain", params.Content)

		dialer := gomail.NewDialer("smtp.gmail.com", 465, myEmail, myPassword)
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		if err := dialer.DialAndSend(message); err != nil {
			return err
		}
	*/
	return nil
}
