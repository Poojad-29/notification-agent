package main

import (
	"log"

	"notification-agent/internal/database"
	"notification-agent/internal/repository"
)

func main() {
	// Connect to Azure PostgreSQL
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	// Create repository
	repo := repository.NotificationRepository{
		DB: db,
	}

	// Save test notification
	err = repo.Save(
		"test@example.com",
		"email",
		"Hello from Azure PostgreSQL",
		"sent",
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Notification saved successfully")
}
