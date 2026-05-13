package server

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/presentation/handlers"
	serverContext "github.com/ggualbertosouza/go-rabbitMq/internal/presentation/server/context"
	"github.com/gin-gonic/gin"
)

func NewRouter(deps serverContext.Dependencies) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/health", handlers.HealthCheck)
	r.POST("/publish", handlers.Publish(deps))

	r.Static("/docs", "./docs")
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/docs")
	})

	return r
}
