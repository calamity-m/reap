package clients

import (
	"fmt"

	"github.com/calamity-m/reap/proto/sow/v1"
	"google.golang.org/grpc"
)

// Returns a connected grpc client to the sow server. Closing of the connection should be handled by the caller.
func CreateSowClient(addr string, opts []grpc.DialOption) (sow.FoodRecordingServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial: %v", err)
	}
	client := sow.NewFoodRecordingServiceClient(conn)
	return client, conn, nil
}
