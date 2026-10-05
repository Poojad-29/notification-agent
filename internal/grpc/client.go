package grpc

import (
	"context"
	"time"

	pb "notification-agent/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClient() error {

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	_, err = client.GetNotificationStatus(
		ctx,
		&pb.StatusRequest{
			NotificationId: "1",
		},
	)

	return err
}
