package routes

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Igwa-Xp/igwa-api-gateway/pkg/auth/pb"
)

type RegisterRequestBody struct {
    Username    string `json:"username"`
    Password string `json:"password"`
	Email string `json:"email"`
	Role string `json:"role"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
    body := RegisterRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Username: body.Username,
        Password: body.Password,
		Email:    body.Email,
		Role: body.Role,
    })

    if err != nil {
        ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    ctx.JSON(int(res.Status), &res)
}