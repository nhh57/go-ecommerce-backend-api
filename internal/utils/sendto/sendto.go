package sendto

import (
	"fmt"
	"github.com/nhh57/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
	"net/smtp"
	"strings"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPORT     = "587"
	SMTPUsername = ""
	SMTPPassword = ""
)

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(email Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg += fmt.Sprintf("From: %s\r\n", email.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(email.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", email.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", email.Body)
	return msg

}

func SentTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Plase enter is to verify your account.", otp),
	}
	messageEmail := BuildMessage(contentEmail)

	// send smtp
	authention := smtp.PlainAuth("", SMTPUsername, SMTPUsername, SMTPHost)

	err := smtp.SendMail(SMTPHost+":25", authention, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Email send faild::", zap.Error(err))
		return err
	}
	return nil
}
