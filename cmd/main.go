package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/auth"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/config"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/experience"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/organizer"
	// "github.com/Igwa-Xp/igwa-api-gateway/pkg/order"
	// "github.com/Igwa-Xp/igwa-api-gateway/pkg/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize authentication service client
	authSvc := auth.InitServiceClient(cfg)

	// Initialize experience service client
	expSvc := experience.InitServiceClient(cfg)

	// Initialize organizer service client
	orgSvc := organizer.InitServiceClient(cfg)

	// Initialize order service client
	// orderSvc := order.InitServiceClient(cfg)

	// Initialize the Gin router
	router := gin.Default()

	// Register routes for each microservice
	auth.RegisterRoutes(router, cfg)
	experience.RegisterRoutes(router, expSvc)
	organizer.RegisterRoutes(router, orgSvc)
	auth.RegisterRoutes(router, cfg, authSvc)

	// Start the server
	log.Printf("Server listening on port %s", cfg.Server.Port)
	err := http.ListenAndServe(":"+cfg.Server.Port, router)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
