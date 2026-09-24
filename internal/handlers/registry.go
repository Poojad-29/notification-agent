package handlers

type Handler interface {
	Send(recipient string, message string) string
}
