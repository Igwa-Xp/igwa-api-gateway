package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/experience"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/experience/pb"
)

func RegisterRoutes(r *gin.Engine, experienceSvc *experience.ServiceClient) {
	v1 := r.Group("/v1")
	experienceRoutes := v1.Group("/experience")

	// Create Experience
	experienceRoutes.POST("/", validateExperienceRequest(), func(c *gin.Context) {
		CreateExperience(c, experienceSvc.Client)
	})

	// Update Experience
	experienceRoutes.PUT("/:id", validateExperienceRequest(), func(c *gin.Context) {
		UpdateExperience(c, experienceSvc.Client)
	})

	// Delete Experience
	experienceRoutes.DELETE("/:id", func(c *gin.Context) {
		DeleteExperience(c, experienceSvc.Client)
	})
}

func validateExperienceRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate request payload using request struct and binding rules
		var request pb.CreateExperienceRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			// Return bad request response if validation fails
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Store the validated request in the context for further processing
		c.Set("experienceRequest", &request)
		c.Next()
	}
}
