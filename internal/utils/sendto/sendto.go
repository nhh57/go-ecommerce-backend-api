package sendto

import (
	"bytes"
	"fmt"
	"github.com/nhh57/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
	"html/template"
	"net/smtp"
	"strings"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPUsername = "datnvpk02264@fpt.edu.vn"
	SMTPPassword = "vvlarxknglgewxvt"
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
	authention := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	err := smtp.SendMail(SMTPHost+":"+SMTPPort, authention, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Email send faild::", zap.Error(err))
		return err
	}
	return nil
}

func SendTemplateEmailOtp(
	to []string, from string, nameTemplate string,
	dataTemplate map[string]interface{},
) error {
	htmlBody, err := getEmailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		return err
	}
	return send(to, from, htmlBody)
}

func getEmailTemplate(
	nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates-html/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}
	messageEmail := BuildMessage(contentEmail)

	// send smtp
	authention := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	err := smtp.SendMail(SMTPHost+":"+SMTPPort, authention, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Email send faild::", zap.Error(err))
		return err
	}
	return nil
}
