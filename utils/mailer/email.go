package mailer

import (
	"bytes"
	"html/template"
	"os"
	"time"

	"github.com/Shabrinashsf/ets-backend-webpro-c.git/config"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/utils"
	"gopkg.in/gomail.v2"
)

const (
	LOCAL_URL          = "http://localhost:3000"
	VERIFY_EMAIL_ROUTE = "register/verify_email"
)

func SendMail(toEmail string, subject string, body string) error {
	emailConfig, err := config.NewEmailConfig()
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailConfig.AuthEmail)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)
	dialer := gomail.NewDialer(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.AuthEmail,
		emailConfig.AuthPassword,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}

func MakeVerificationEmail(receiverEmail string) (map[string]string, error) {
	expired := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	plainText := receiverEmail + "_" + expired
	token, err := utils.AESEncrypt(plainText)
	if err != nil {
		return nil, err
	}

	verifyLink := LOCAL_URL + "/" + VERIFY_EMAIL_ROUTE + "?token=" + token

	readHtml, err := os.ReadFile("utils/email-template/base_mail.html")
	if err != nil {
		return nil, err
	}

	data := struct {
		Email  string
		Verify string
	}{
		Email:  receiverEmail,
		Verify: verifyLink,
	}

	tmpl, err := template.New("custom").Parse(string(readHtml))
	if err != nil {
		return nil, err
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		return nil, err
	}

	draftEmail := map[string]string{
		"subject": "Cakno - Go Gin Template",
		"body":    strMail.String(),
	}

	return draftEmail, nil
}
