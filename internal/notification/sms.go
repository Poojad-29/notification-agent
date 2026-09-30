package notification

import (
	"os"

	twilio "github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(to string, body string) error {
	client := twilio.NewRestClientWithParams(
		twilio.ClientParams{
			Username: os.Getenv("TWILIO_SID"),
			Password: os.Getenv("TWILIO_TOKEN"),
		},
	)

	params := &api.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(os.Getenv("TWILIO_PHONE"))
	params.SetBody(body)

	_, err := client.Api.CreateMessage(params)
	return err
}
