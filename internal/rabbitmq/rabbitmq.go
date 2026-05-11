package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func New(url string) *RabbitMq {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return &RabbitMq{
		Connection: conn,
		Channel:    ch,
	}
}

func (r *RabbitMq) Setup(exchange, queue, routingKey string) error {
	if err := r.DeclareExchange(exchange); err != nil {
		return err
	}

	if err := r.DeclareQueue(queue); err != nil {
		return err
	}

	if err := r.BindQueue(queue, routingKey, exchange); err != nil {
		return err
	}

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
