package repository

type NotificationRepository interface {
	Save() error
	GetByID(id string) error
}
