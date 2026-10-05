package main

import (
	"log"

	"github.com/joho/godotenv"

	"notification-agent/internal/database"
	grpcservice "notification-agent/internal/grpc"
	"notification-agent/internal/notification"
	"notification-agent/internal/repository"
	"notification-agent/internal/servicebus"
	"notification-agent/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// SendGrid Email Test
	err = notification.SendEmail(
		"pooja.b.d@capgemini.com",
		"Notification Agent Test",
		"Hello from SendGrid Email Integration",
	)

	if err != nil {
		log.Println("Email Error:", err)
	} else {
		log.Println("Email sent successfully")
	}

	// Twilio SMS Test
	err = notification.SendSMS(
		"+918073550593",
		"Hello from Notification Agent",
	)

	if err != nil {
		log.Println("SMS Error:", err)
	} else {
		log.Println("SMS sent successfully")
	}

	// Database Connection
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	// Repository
	repo := repository.NotificationRepository{
		DB: db,
	}

	// Service Layer
	service := services.NotificationService{
		Repo: &repo,
	}

	// gRPC Service
	grpcSvc := grpcservice.NotificationGRPCService{
		Service: &service,
	}

	// Create Notification
	response, err := grpcSvc.CreateNotification(
		"test@example.com",
		"email",
		"Hello from gRPC Service",
	)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.Status, response.Message)
	}

	// Get Notifications
	rows, err := service.GetNotifications()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var recipient string
		var channel string
		var message string
		var status string
		var createdAt string

		err := rows.Scan(
			&id,
			&recipient,
			&channel,
			&message,
			&status,
			&createdAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf(
			"ID=%d Recipient=%s Channel=%s Message=%s Status=%s CreatedAt=%s",
			id,
			recipient,
			channel,
			message,
			status,
			createdAt,
		)
	}

	log.Println("Notifications retrieved successfully")

	// Azure Service Bus Consumer Test
	err = servicebus.ReceiveMessage()
	if err != nil {
		log.Println("Service Bus Receive Error:", err)
	} else {
		log.Println("Service Bus message processed successfully")
	}
}
