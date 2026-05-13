package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMq) Publish(
	exchange AppExchanges,
	routingKey AppRoutingKeys,
	body []byte,
) error {
	r.Logger.Info("publishing message")

	err := r.Channel.Publish(
		string(exchange),
		string(routingKey),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
