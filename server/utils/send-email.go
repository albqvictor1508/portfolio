package utils

import (
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"
)

type SendEmailParams struct {
	to      string
	subject *string
	content string
}

func NewMailer(to string, subject *string, content string) SendEmailParams {
	return SendEmailParams{
		to:      to,
		subject: subject,
		content: content,
	}
}

func SendEmail(to string) error {
	var emailParams *SendEmailParams
	myEmail := os.Getenv("MY_EMAIL")
	myPassword := os.Getenv("MY_PASSWORD")

	message := gomail.NewMessage()

	message.SetHeader("From", fmt.Sprintf("from <%v>", myEmail))
	message.SetHeader("To", emailParams.to)
	message.SetHeader("Subject", *emailParams.subject)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}
