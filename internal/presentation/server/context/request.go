package serverContext

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bind[T any](ctx *gin.Context) (*T, bool) {
	var req T

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return nil, false
	}

	return &req, true
}
