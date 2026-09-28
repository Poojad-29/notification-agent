package services

import (
	"database/sql"

	"notification-agent/internal/repository"
)

type NotificationService struct {
	Repo *repository.NotificationRepository
}

func (s *NotificationService) CreateNotification(
	recipient string,
	channel string,
	message string,
) error {

	return s.Repo.Save(
		recipient,
		channel,
		message,
		"sent",
	)
}

func (s *NotificationService) GetNotifications() (*sql.Rows, error) {
	return s.Repo.GetAll()
}

func (s *NotificationService) UpdateNotificationStatus(
	id int,
	status string,
) error {

	return s.Repo.UpdateStatus(
		id,
		status,
	)
}

func (s *NotificationService) DeleteNotification(
	id int,
) error {

	return s.Repo.Delete(id)
}
