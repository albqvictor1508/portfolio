package utils

import (
	"crypto/tls"
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"
)

type SendEmailParams struct {
	ReplyTo string `json:"reply_to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

var (
	myEmail    string = os.Getenv("MY_EMAIL")
	myPassword string = os.Getenv("MY_PASSWORD")
)

func SendEmail(params SendEmailParams) error {
	message := gomail.NewMessage()

	fmt.Print(params)

	if params.Subject == "" {
		params.Subject = "By Portfolio"
	}

	if params.ReplyTo == "" {
		return fmt.Errorf("ReplyTo is required")
	}

	message.SetHeader("From", myEmail)
	message.SetHeader("Reply-To", params.ReplyTo)
	message.SetHeader("To", myEmail)
	message.SetHeader("Subject", params.Subject)
	message.SetBody("text/plain", params.Content)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("error to send message: %v", err)
	}

	return nil
}
