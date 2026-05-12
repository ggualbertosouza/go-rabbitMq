package handlers

import (
	"net/http"

	"github.com/ggualbertosouza/go-rabbitMq/internal/server/context"
	"github.com/gin-gonic/gin"
)

func Publish(deps context.Dependencies) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		deps.Logger.Info("publishing message")

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
