package services

import (
	"backend/internal/config"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type EmailService interface {
	SendEmail(to, subject, body string) error
}

type emailService struct {
	config *config.Config
}

func NewEmailService(cfg *config.Config) EmailService {
	return &emailService{config: cfg}
}

func (e *emailService) SendEmail(to, subject, body string) error {

	if strings.HasSuffix(to, "@mock.com") {
		log.Println("Email not sent to mock email address")
		return nil
	}

	auth := smtp.PlainAuth("", e.config.SMTP.Username, e.config.SMTP.Password, e.config.SMTP.Host)

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))
	addr := fmt.Sprintf("%s:%d", e.config.SMTP.Host, e.config.SMTP.Port)

	return smtp.SendMail(addr, auth, e.config.SMTP.From, []string{to}, msg)
}
