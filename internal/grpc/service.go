package grpc

import (
	"notification-agent/internal/services"
)

type NotificationGRPCService struct {
	Service *services.NotificationService
}

func (g *NotificationGRPCService) CreateNotification(
	recipient string,
	channel string,
	message string,
) (*StatusResponse, error) {

	err := g.Service.CreateNotification(
		recipient,
		channel,
		message,
	)

	if err != nil {
		return &StatusResponse{
			Status:  "FAILED",
			Message: err.Error(),
		}, err
	}

	return &StatusResponse{
		Status:  "SUCCESS",
		Message: "Notification created successfully",
	}, nil
}
