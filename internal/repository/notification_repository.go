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

func (r *NotificationRepository) GetAll() (*sql.Rows, error) {
	query := `
	SELECT id, recipient, channel, message, status, created_at
	FROM notifications
	ORDER BY id
	`

	return r.DB.Query(query)
}

func (r *NotificationRepository) UpdateStatus(id int, status string) error {
	query := `
	UPDATE notifications
	SET status = $1
	WHERE id = $2
	`

	_, err := r.DB.Exec(query, status, id)
	return err
}

func (r *NotificationRepository) Delete(id int) error {
	query := `
	DELETE FROM notifications
	WHERE id = $1
	`

	_, err := r.DB.Exec(query, id)
	return err
}
