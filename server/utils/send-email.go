package utils

import (
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"
)

type SendEmailParams struct {
	To      string  `json:"user_email"`
	Subject *string `json:"subject"`
	Content string  `json:"content"`
}

func SendEmail(params SendEmailParams) error {
	myEmail := os.Getenv("MY_EMAIL")
	myPassword := os.Getenv("MY_PASSWORD")

	message := gomail.NewMessage()

	message.SetHeader("From", fmt.Sprintf("from <%v>", myEmail))
	message.SetHeader("To", params.To)
	message.SetHeader("Subject", *params.Subject)
	message.SetBody("text/plain", params.Content)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}
