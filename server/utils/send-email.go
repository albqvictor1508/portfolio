package utils

import (
	"os"

	gomail "gopkg.in/mail.v2"
)

type SendEmailParams struct {
	From    string  `json:"user_email"`
	Subject *string `json:"subject"`
	Content string  `json:"content"`
}

var (
	myEmail    string = os.Getenv("MY_EMAIL")
	myPassword string = os.Getenv("MY_PASSWORD")
)

func SendEmail(params SendEmailParams) error {
	message := gomail.NewMessage()

	message.SetHeader("From", params.From)
	message.SetHeader("To", myEmail)
	message.SetHeader("Subject", *params.Subject)
	message.SetBody("text/plain", params.Content)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}
