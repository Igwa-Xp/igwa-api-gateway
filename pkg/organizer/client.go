package organizer

import (
	"log"
	"github.com/Igwa-Xp/igwa-api-gateway/config"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/organizer/pb"
	"google.golang.org/grpc"
	
)

type ServiceClient struct {
	Client pb.OrganizerServiceClient
}

func InitServiceClient(c *config.Config) pb.OrganizerServiceClient {
	conn, err := grpc.Dial(c.OrganizerServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the organizer service: %v", err)
	}

	client := pb.NewOrganizerServiceClient(conn)

	return client
}


// Implement any additional helper methods or initialization logic as needed
