package handlers

import (
	"net/http"

	userUseCase "github.com/ggualbertosouza/go-rabbitMq/internal/application/useCases/user"
	serverContext "github.com/ggualbertosouza/go-rabbitMq/internal/presentation/server/context"
	"github.com/gin-gonic/gin"
)

type publishRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Publish(deps serverContext.Dependencies) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, ok := serverContext.Bind[publishRequest](ctx)
		if !ok {
			return
		}

		userUseCase := userUseCase.New(deps.Logger, deps.Rabbit)
		if err := userUseCase.CreateUser(req.Email, req.Password); err != nil {
			serverContext.BadRequest(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
