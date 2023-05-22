package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/auth/routes"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/config"
)

const (
	apiVersion = "v1"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	authGroup := r.Group("/auth/" + apiVersion)

	authGroup.POST("/register", svc.Register)
	authGroup.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
