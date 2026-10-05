package servicebus

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func ReceiveMessage() error {
	client, err := azservicebus.NewClientFromConnectionString(
		os.Getenv("SERVICE_BUS_CONNECTION_STRING"),
		nil,
	)
	if err != nil {
		return err
	}
	defer client.Close(context.Background())

	receiver, err := client.NewReceiverForQueue(
		os.Getenv("SERVICE_BUS_QUEUE_NAME"),
		nil,
	)
	if err != nil {
		return err
	}

	messages, err := receiver.ReceiveMessages(
		context.Background(),
		1,
		nil,
	)
	if err != nil {
		return err
	}

	for _, message := range messages {

		ProcessMessage(string(message.Body))

		err = receiver.CompleteMessage(
			context.Background(),
			message,
			nil,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
