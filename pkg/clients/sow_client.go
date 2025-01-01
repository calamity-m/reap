package clients

import (
	"fmt"

	"github.com/calamity-m/reap/proto/sow/v1"
	"google.golang.org/grpc"
)

func CreateSowClient(addr string, opts []grpc.DialOption) (sow.FoodRecordingServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial: %v", err)
	}
	client := sow.NewFoodRecordingServiceClient(conn)
	return client, conn, nil
}
