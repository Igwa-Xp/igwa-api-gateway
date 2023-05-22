package routes

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Igwa-Xp/igwa-api-gateway/pkg/organizer/pb"
)

func DeleteOrganizer(ctx *gin.Context, c pb.OrganizerServiceClient) {
    organizerID := ctx.Param("id")

    res, err := c.DeleteOrganizer(context.Background(), &pb.DeleteOrganizerRequest{
        Id: organizerID,
    })

    if err != nil {
        ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    ctx.JSON(int(res.Status), &res)
}
