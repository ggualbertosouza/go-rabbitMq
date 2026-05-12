package rabbitmq

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/config/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	Logger     logger.Logger
	Channel    *amqp.Channel
	Connection *amqp.Connection
}

func New(log logger.Logger, url string) (*RabbitMq, error) {
	log.Info("connecting to rabbitmq", logger.String("url", url))

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Error("failed to connect rabbitmq", logger.Error(err))
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Error("failed to create rabbitmq channel", logger.Error(err))
		return nil, err
	}

	log.Info("rabbitmq connected")

	return &RabbitMq{
		Connection: conn,
		Channel:    ch,
		Logger:     log,
	}, nil
}
