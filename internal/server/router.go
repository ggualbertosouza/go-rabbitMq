package server

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/handlers"
	"github.com/ggualbertosouza/go-rabbitMq/internal/server/context"
	"github.com/gin-gonic/gin"
)

func NewRouter(deps context.Dependencies) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/health", handlers.HealthCheck)
	r.POST("/publish", handlers.Publish(deps))

	return r
}
