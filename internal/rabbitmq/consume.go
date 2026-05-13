package rabbitmq

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/config/logger"
)

func (r *RabbitMq) Consume(
	queue string,
	consumer string,
) error {
	r.Logger.Info("starting consuming")

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
		return err
	}

	go func() {
		for message := range messages {
			r.Logger.Info("message received", logger.String("body", string(message.Body)))

			err := message.Ack(false)
			if err != nil {
				r.Logger.Error("failed  to ack message", logger.Error(err))
			}
		}
	}()

	return nil
}
