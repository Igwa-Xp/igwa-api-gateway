package routes

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Igwa-Xp/igwa-api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
    b := LoginRequestBody{}

    if err := ctx.BindJSON(&b); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    res, err := c.Login(context.Background(), &pb.LoginRequest{
        Username: b.Username,
        Password: b.Password,
    })

    if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    ctx.JSON(int(res.Status), &res)
}