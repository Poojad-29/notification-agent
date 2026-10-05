package grpc

import (
	"context"

	pb "notification-agent/proto"
)

type GRPCServer struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *GRPCServer) CreateNotification(
	ctx context.Context,
	req *pb.NotificationRequest,
) (*pb.NotificationResponse, error) {

	return &pb.NotificationResponse{
		Status:  "SUCCESS",
		Message: "Notification created successfully",
	}, nil
}

func (s *GRPCServer) GetNotificationStatus(
	ctx context.Context,
	req *pb.StatusRequest,
) (*pb.StatusResponse, error) {

	return &pb.StatusResponse{
		Status: "sent",
	}, nil
}
