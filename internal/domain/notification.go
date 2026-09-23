package domain

type Notification struct {
	ID        string
	Recipient string
	Channel   string
	Message   string
	Status    string
}
