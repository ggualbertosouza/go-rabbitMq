package rabbitmq

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/config/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Logger     logger.Logger
}

func New(log logger.Logger, url string) (*RabbitMq, error) {
	log.Info(
		"connecting to rabbitmq",
		logger.String("url", url),
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Error(
			"failed to connect rabbitmq",
			logger.Any("error", err),
		)

		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Error(
			"failed to create rabbitmq channel",
			logger.Any("error", err),
		)

		return nil, err
	}

	log.Info("rabbitmq connected")

	return &RabbitMq{
		Connection: conn,
		Channel:    ch,
		Logger:     log,
	}, nil
}

func (r *RabbitMq) Init(exchange, queue, routingKey string) error {
	r.Logger.Info(
		"initializing rabbitmq topology",
		logger.String("exchange", exchange),
		logger.String("queue", queue),
		logger.String("routing_key", routingKey),
	)

	if err := r.DeclareExchange(exchange); err != nil {
		r.Logger.Error(
			"failed to declare exchange",
			logger.Any("error", err),
			logger.String("exchange", exchange),
		)

		return err
	}

	if err := r.DeclareQueue(queue); err != nil {
		r.Logger.Error(
			"failed to declare queue",
			logger.Any("error", err),
			logger.String("queue", queue),
		)

		return err
	}

	if err := r.BindQueue(queue, routingKey, exchange); err != nil {
		r.Logger.Error(
			"failed to bind queue",
			logger.Any("error", err),
			logger.String("queue", queue),
			logger.String("exchange", exchange),
			logger.String("routing_key", routingKey),
		)

		return err
	}

	r.Logger.Info(
		"rabbitmq topology initialized",
		logger.String("exchange", exchange),
		logger.String("queue", queue),
	)

	return nil
}

func (r *RabbitMq) DeclareExchange(name string) error {
	return r.Channel.ExchangeDeclare(name, "direct", true, false, false, false, nil)
}

func (r *RabbitMq) DeclareQueue(name string) error {
	_, err := r.Channel.QueueDeclare(name, true, false, false, false, nil)
	return err
}

func (r *RabbitMq) BindQueue(queue, routingKey, exchange string) error {
	return r.Channel.QueueBind(queue, routingKey, exchange, false, nil)
}
