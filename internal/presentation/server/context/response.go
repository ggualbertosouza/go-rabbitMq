package serverContext

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(ctx *gin.Context, status int, data any) {
	if data == nil {
		ctx.Status(status)
		return
	}

	ctx.JSON(status, data)
}

func Ok(ctx *gin.Context, data any) {
	JSON(ctx, http.StatusOK, data)
}

func BadRequest(ctx *gin.Context, err error) {
	JSON(ctx, http.StatusBadRequest, err.Error())
}
