package notification

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(to string, subject string, body string) error {
	from := mail.NewEmail(
		"Notification Agent",
		os.Getenv("SENDGRID_FROM_EMAIL"),
	)

	recipient := mail.NewEmail("", to)

	message := mail.NewSingleEmail(
		from,
		subject,
		recipient,
		body,
		body,
	)

	client := sendgrid.NewSendClient(
		os.Getenv("SENDGRID_API_KEY"),
	)

	_, err := client.Send(message)

	return err
}
