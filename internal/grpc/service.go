package grpc

import (
	"notification-agent/internal/services"
)

type NotificationGRPCService struct {
	Service *services.NotificationService
}
