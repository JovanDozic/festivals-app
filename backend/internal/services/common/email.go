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
	SendOrderConfirmation(to, subject, body, orderType string, order models.Order) error
	SendBraceletIssuedConfirmation(to, subject, body string, bracelet models.Bracelet) error
	SendBraceletActivatedConfirmation(to, subject, body string, bracelet models.Bracelet) error
	SendBraceletHelpRequestApproved(to, subject, body string, bracelet models.Bracelet) error
	SendBraceletHelpRequestRejected(to, subject, body string, bracelet models.Bracelet) error
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

func (e *emailService) SendOrderConfirmation(to, subject, body, orderType string, order models.Order) error {

	templateData, err := os.ReadFile("assets/template.html")
	if err != nil {
		return fmt.Errorf("failed to read email template: %w", err)
	}

	template := string(templateData)

	template = strings.ReplaceAll(template, "#TITLE#", fmt.Sprintf("%s Order Confirmation", orderType))
	template = strings.ReplaceAll(template, "#MESSAGE#", fmt.Sprintf(`Order (ID: <span class="number-mono">%d</span>) has been successfully placed.`, order.ID))
	template = strings.ReplaceAll(template, "#CONTENT#", fmt.Sprintf(`Total amount: <span class="number-mono">$%.2f</span> <br>`, order.TotalAmount))
	template = strings.ReplaceAll(template, "#BOTTOMMESSAGE#", "We will notify you via email once your Bracelet is shipped.")
	template = strings.ReplaceAll(template, "#YEAR#", time.Now().Format("2006"))

	if err := e.SendEmail(to, subject, template); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil

}

func (e *emailService) SendBraceletIssuedConfirmation(to, subject, body string, bracelet models.Bracelet) error {

	templateData, err := os.ReadFile("assets/template.html")
	if err != nil {
		return fmt.Errorf("failed to read email template: %w", err)
	}

	template := string(templateData)

	template = strings.ReplaceAll(template, "#TITLE#", "Bracelet Issued Confirmation")
	template = strings.ReplaceAll(template, "#MESSAGE#", fmt.Sprintf(`Bracelet (barcode: <span class="number-mono">%s</span>) has been successfully issued.`, bracelet.BarcodeNumber))
	template = strings.ReplaceAll(template, "#CONTENT#", "When you receive your bracelet, please activate it by going to My Orders page.")
	template = strings.ReplaceAll(template, "#BOTTOMMESSAGE#", "")
	template = strings.ReplaceAll(template, "#YEAR#", time.Now().Format("2006"))

	if err := e.SendEmail(to, subject, template); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (e *emailService) SendBraceletActivatedConfirmation(to, subject, body string, bracelet models.Bracelet) error {

	templateData, err := os.ReadFile("assets/template.html")
	if err != nil {
		return fmt.Errorf("failed to read email template: %w", err)
	}

	template := string(templateData)

	template = strings.ReplaceAll(template, "#TITLE#", "Bracelet Activated Confirmation")
	template = strings.ReplaceAll(template, "#MESSAGE#", fmt.Sprintf(`Bracelet (barcode: <span class="number-mono">%s</span>) has been successfully activated.`, bracelet.BarcodeNumber))
	template = strings.ReplaceAll(template, "#CONTENT#", "You can now use your bracelet to make purchases at the festival. To top-up your bracelet, please go to My Bracelets page.")
	template = strings.ReplaceAll(template, "#BOTTOMMESSAGE#", "")
	template = strings.ReplaceAll(template, "#YEAR#", time.Now().Format("2006"))

	if err := e.SendEmail(to, subject, template); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (e *emailService) SendBraceletHelpRequestApproved(to, subject, body string, bracelet models.Bracelet) error {

	templateData, err := os.ReadFile("assets/template.html")
	if err != nil {
		return fmt.Errorf("failed to read email template: %w", err)
	}

	template := string(templateData)

	template = strings.ReplaceAll(template, "#TITLE#", "Help Request Approved")
	template = strings.ReplaceAll(template, "#MESSAGE#", fmt.Sprintf(`Help request for bracelet (barcode: <span class="number-mono">%s</span>) has been approved.`, bracelet.BarcodeNumber))
	template = strings.ReplaceAll(template, "#CONTENT#", "You can now use your bracelet to make purchases at the festival. To top-up your bracelet, please go to My Bracelets page.")
	template = strings.ReplaceAll(template, "#BOTTOMMESSAGE#", "")
	template = strings.ReplaceAll(template, "#YEAR#", time.Now().Format("2006"))

	if err := e.SendEmail(to, subject, template); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (e *emailService) SendBraceletHelpRequestRejected(to, subject, body string, bracelet models.Bracelet) error {

	templateData, err := os.ReadFile("assets/template.html")
	if err != nil {
		return fmt.Errorf("failed to read email template: %w", err)
	}

	template := string(templateData)

	template = strings.ReplaceAll(template, "#TITLE#", "Help Request Rejected")
	template = strings.ReplaceAll(template, "#MESSAGE#", fmt.Sprintf(`Help request for bracelet (barcode: <span class="number-mono">%s</span>) has been rejected.`, bracelet.BarcodeNumber))
	template = strings.ReplaceAll(template, "#CONTENT#", "If you think we made a mistake, please contact the festival staff or customer support for further assistance.")
	template = strings.ReplaceAll(template, "#BOTTOMMESSAGE#", "")
	template = strings.ReplaceAll(template, "#YEAR#", time.Now().Format("2006"))

	if err := e.SendEmail(to, subject, template); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
