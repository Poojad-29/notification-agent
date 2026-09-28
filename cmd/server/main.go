package main

import (
	"log"

	"notification-agent/internal/database"
	"notification-agent/internal/repository"
	"notification-agent/internal/services"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	repo := repository.NotificationRepository{
		DB: db,
	}

	service := services.NotificationService{
		Repo: &repo,
	}

	err = service.CreateNotification(
		"test@example.com",
		"email",
		"Hello from Notification Service",
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Notification created successfully")

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

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Notifications retrieved successfully")
}
