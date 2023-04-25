package newEmail

import (
	message "send-email/internal"
	"send-email/internal/email"
)

func SendHandler(summary message.Summary, enviarService *email.EmailService) error {
	err := enviarService.Send(summary)
	if err != nil {
		return err
	}
	return nil
}
