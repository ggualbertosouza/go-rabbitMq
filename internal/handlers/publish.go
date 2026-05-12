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

		l.Info("Starting publishing message")

		var req publishRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		err := deps.Rabbit.Publish(
			"events",
			"demo.message",
			[]byte(req.Message),
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
