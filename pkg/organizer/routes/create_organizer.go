package routes

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Igwa-Xp/igwa-api-gateway/pkg/organizer/pb"
)

type CreateOrganizerRequestBody struct {
    Name            string   `json:"name"`
    Email           string   `json:"email"`
    PhoneNumbers    []string `json:"phoneNumbers"`
    OrganizerType   string   `json:"organizerType"`
    KYCDocumentType string   `json:"kycDocumentType"`
    KYCDocumentNumber string `json:"kycDocumentNumber"`
    // Add any additional KYC-related fields as needed
}

func CreateOrganizer(ctx *gin.Context, c pb.OrganizerServiceClient) {
    body := CreateOrganizerRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    res, err := c.CreateOrganizer(context.Background(), &pb.CreateOrganizerRequest{
        Name:            body.Name,
        Email:           body.Email,
        PhoneNumbers:    body.PhoneNumbers,
        OrganizerType:   body.OrganizerType,
        KYCDocumentType: body.KYCDocumentType,
        KYCDocumentNumber: body.KYCDocumentNumber,
    })

    if err != nil {
        ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    ctx.JSON(int(res.Status), &res)
}
