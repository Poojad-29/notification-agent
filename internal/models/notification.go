package models

type Notification struct {
	ID        int
	Recipient string
	Channel   string
	Message   string
	Status    string
	CreatedAt string
}
