package servicebus

import "fmt"

func SendMessage(message string) error {
	fmt.Println("Service Bus Producer:", message)
	return nil
}
