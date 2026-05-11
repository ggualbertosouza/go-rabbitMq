package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Publish(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"published": "ok",
	})
}
