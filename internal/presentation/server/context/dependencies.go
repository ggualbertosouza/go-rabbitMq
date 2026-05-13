package serverContext

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/rabbitmq"
)

type Dependencies struct {
	Logger logger.Logger
	Rabbit *rabbitmq.RabbitMq
}
