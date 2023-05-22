package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/organizer"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/organizer/pb"
)

func RegisterRoutes(r *gin.Engine, organizerSvc *organizer.ServiceClient) {
	v1Routes := r.Group("/v1")

	organizerRoutes := v1Routes.Group("/organizer")

	// Create Organizer
	organizerRoutes.POST("/", func(c *gin.Context) {
		CreateOrganizer(c, organizerSvc.Client)
	})

	// Update Organizer
	organizerRoutes.PUT("/:id", func(c *gin.Context) {
		UpdateOrganizer(c, organizerSvc.Client)
	})

	// Delete Organizer
	organizerRoutes.DELETE("/:id", func(c *gin.Context) {
		DeleteOrganizer(c, organizerSvc.Client)
	})
}
