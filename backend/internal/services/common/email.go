package common

import (
	"backend/internal/config"
	models "backend/internal/models/festival"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type EmailService interface {
	SendEmail(to, subject, body string) error
	SendBraceletTopUpConfirmation(to, subject, body string, bracelet models.Bracelet, amount float64) error
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

	go func() {
		auth := smtp.PlainAuth("", e.config.SMTP.Username, e.config.SMTP.Password, e.config.SMTP.Host)

		// Build the email message with MIME headers for HTML content
		msg := []byte(fmt.Sprintf(
			"To: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
			to, subject, body,
		))

		addr := fmt.Sprintf("%s:%d", e.config.SMTP.Host, e.config.SMTP.Port)

		if err := smtp.SendMail(addr, auth, e.config.SMTP.From, []string{to}, msg); err != nil {
			log.Printf("Error sending email: %v", err)
		}
	}()
	return nil
}

func (e *emailService) SendBraceletTopUpConfirmation(to, subject, body string, bracelet models.Bracelet, amount float64) error {

	templateData, err := os.ReadFile("assets/template.html")
	if err != nil {
		return fmt.Errorf("failed to read email template: %w", err)
	}

	template := string(templateData)

	template = strings.ReplaceAll(template, "#TITLE#", "Top-up Confirmation")
	template = strings.ReplaceAll(template, "#MESSAGE#", fmt.Sprintf(`Top-up for bracelet (barcode: <span class="number-mono">%s</span>) has been successful.`, bracelet.BarcodeNumber))
	template = strings.ReplaceAll(template, "#CONTENT#", fmt.Sprintf(`Top-up amount: <span class="number-mono">$%.2f</span> <br> Current balance: <span class="number-mono"><b>$%.2f</b></span>`, amount, bracelet.Balance))
	template = strings.ReplaceAll(template, "#BOTTOMMESSAGE#", "")
	template = strings.ReplaceAll(template, "#YEAR#", time.Now().Format("2006"))

	if err := e.SendEmail(to, subject, template); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
