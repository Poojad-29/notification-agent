package servicebus

import "fmt"

func ProcessMessage(message string) {
	fmt.Println("Service Bus Consumer:", message)
	fmt.Println("Notification processed successfully")
}
