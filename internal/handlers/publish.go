package handlers

import (
	"net/http"

	"github.com/ggualbertosouza/go-rabbitMq/internal/server/context"
	"github.com/gin-gonic/gin"
)

type publishRequest struct {
	Message string `json:"message"`
}

func Publish(deps context.Dependencies) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		l := deps.Logger

		l.Info("starting publishing message")

		req, ok := context.Bind[publishRequest](ctx)
		if !ok {
			return
		}

		l.Info(req.Message)

		l.Info("message sent successfully")
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
