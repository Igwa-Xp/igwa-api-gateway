package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/experience/pb"
)

type UpdateExperienceRequestBody struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Description string   `json:"description"`
	Categories []string `json:"categories"`
	OrganizerID string   `json:"organizerId"`
	Images     []string `json:"images"`
	AltImage   string   `json:"altImage"`
	Price      float64  `json:"price"`
	Duration   int32    `json:"duration"`
	Location   string   `json:"location"`
	Category   string   `json:"category"`
	Tags       []string `json:"tags"`
	// Add any additional fields as needed
}

func UpdateExperience(ctx *gin.Context, c pb.ExperienceServiceClient) {
	body := UpdateExperienceRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateExperience(context.Background(), &pb.UpdateExperienceRequest{
		Id:          body.ID,
		Title:       body.Title,
		Description: body.Description,
		Categories:  body.Categories,
		OrganizerId: body.OrganizerID,
		Images:      body.Images,
		AltImage:    body.AltImage,
		Price:       body.Price,
		Duration:    body.Duration,
		Location:    body.Location,
		Category:    body.Category,
		Tags:        body.Tags,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
