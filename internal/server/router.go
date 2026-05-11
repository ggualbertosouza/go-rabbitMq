package server

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/health", handlers.HealthCheck)
	r.POST("/publish", handlers.Publish)

	return r
}
