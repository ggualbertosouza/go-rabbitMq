package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMq) Publish(exchange ExchangeType, routingKey string, body []byte) error {
	r.Logger.Info("publishing message")

	return r.Channel.PublishWithContext(
		context.Background(),
		string(exchange),
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
