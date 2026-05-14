package rabbitmq

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	"github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMq) Consume(
	queue string,
	consumer string,
) (<-chan amqp091.Delivery, error) {
	r.Logger.Info("starting consuming", logger.String("Consumer", consumer))

	messages, err := r.Channel.Consume(
		queue,
		consumer,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return messages, nil
}
