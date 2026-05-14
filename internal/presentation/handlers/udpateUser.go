package handlers

import (
	"net/http"

	userUseCase "github.com/ggualbertosouza/go-rabbitMq/internal/application/useCases/user"
	serverContext "github.com/ggualbertosouza/go-rabbitMq/internal/presentation/server/context"
	"github.com/gin-gonic/gin"
)

type updateUserRequest struct {
	Email string `json:"email" binding:"required"`
}

func UpdateUser(deps serverContext.Dependencies) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, ok := serverContext.Bind[updateUserRequest](ctx)
		if !ok {
			return
		}

		userUseCase := userUseCase.New(deps.Logger, deps.Rabbit)
		if err := userUseCase.Update(req.Email); err != nil {
			serverContext.BadRequest(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
