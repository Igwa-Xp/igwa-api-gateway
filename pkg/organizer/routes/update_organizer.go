package routes

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Igwa-Xp/igwa-api-gateway/pkg/organizer/pb"
)

type UpdateOrganizerRequestBody struct {
    ID            string   `json:"id"`
    Name          string   `json:"name"`
    Email         string   `json:"email"`
    PhoneNumbers  []string `json:"phoneNumbers"`
    // Add any additional fields as needed
}

func UpdateOrganizer(ctx *gin.Context, c pb.OrganizerServiceClient) {
    body := UpdateOrganizerRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    res, err := c.UpdateOrganizer(context.Background(), &pb.UpdateOrganizerRequest{
        Id:            body.ID,
        Name:          body.Name,
        Email:         body.Email,
        PhoneNumbers:  body.PhoneNumbers,
        // Add any additional fields as needed
    })

    if err != nil {
        ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    ctx.JSON(int(res.Status), &res)
}
