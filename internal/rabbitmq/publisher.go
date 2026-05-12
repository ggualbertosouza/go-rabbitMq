package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMq) Publish(exchange, routingKey string, body []byte) error {
	r.Logger.Info("publishing message")

	return r.Channel.PublishWithContext(
		context.Background(),
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
