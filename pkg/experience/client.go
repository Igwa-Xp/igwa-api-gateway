package experience

import (
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/config"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/experience/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ExperienceServiceClient
}

func InitServiceClient(c *config.Config) pb.ExperienceServiceClient {
	conn, err := grpc.Dial(c.ExperienceServiceURL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return pb.NewExperienceServiceClient(conn)
}
