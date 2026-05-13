package userRbMq

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type UserRabbitMq struct {
	Logger  logger.Logger
	channel *amqp.Channel
}

func New(log logger.Logger, ch *amqp.Channel) *UserRabbitMq {
	return &UserRabbitMq{
		Logger:  log,
		channel: ch,
	}
}

func (u *UserRabbitMq) Publish(
	exchange AppExchanges,
	routingKey AppRoutingKeys,
	body []byte,
) error {
	u.Logger.Info("publishing message")

	err := u.channel.Publish(
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
		u.Logger.Info("error while publishing message")
		return err
	}

	u.Logger.Info("message published successfully")
	return nil
}
