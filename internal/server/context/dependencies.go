package context

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/config/logger"
	"github.com/ggualbertosouza/go-rabbitMq/internal/rabbitmq"
)

type Dependencies struct {
	Logger logger.Logger
	Rabbit *rabbitmq.RabbitMq
}
