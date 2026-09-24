package notification

type NotificationEvent struct {
	ID        string
	Channel   string
	Recipient string
	Message   string
}
