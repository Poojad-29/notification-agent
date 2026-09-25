package repository

import "database/sql"

type NotificationRepository struct {
	DB *sql.DB
}

func (r *NotificationRepository) Save(
	recipient string,
	channel string,
	message string,
	status string,
) error {

	query := `
	INSERT INTO notifications
	(recipient, channel, message, status)
	VALUES ($1, $2, $3, $4)
	`

	_, err := r.DB.Exec(
		query,
		recipient,
		channel,
		message,
		status,
	)

	return err
}
