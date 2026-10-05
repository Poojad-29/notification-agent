package servicebus

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func CreateClient() (*azservicebus.Client, error) {
	return azservicebus.NewClientFromConnectionString(
		os.Getenv("SERVICE_BUS_CONNECTION_STRING"),
		nil,
	)
}

func SendAzureMessage(message string) error {
	client, err := CreateClient()
	if err != nil {
		return err
	}
	defer client.Close(context.Background())

	sender, err := client.NewSender(
		os.Getenv("SERVICE_BUS_QUEUE_NAME"),
		nil,
	)
	if err != nil {
		return err
	}

	return sender.SendMessage(
		context.Background(),
		&azservicebus.Message{
			Body: []byte(message),
		},
		nil,
	)
}
