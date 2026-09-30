package servicebus

import "fmt"

func ReceiveMessage(message string) {
	fmt.Println("Service Bus Consumer:", message)

	fmt.Println("Notification processed successfully")
}
