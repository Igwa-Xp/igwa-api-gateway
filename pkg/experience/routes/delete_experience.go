package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Igwa-Xp/igwa-api-gateway/pkg/experience/pb"
)

func DeleteExperience(ctx *gin.Context, c pb.ExperienceServiceClient) {
	experienceID := ctx.Param("id")

	res, err := c.DeleteExperience(context.Background(), &pb.DeleteExperienceRequest{
		Id: experienceID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
