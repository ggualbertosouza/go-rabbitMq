package rabbitmq

func (r *RabbitMq) Publish(exchange ExchangeType, body []byte) error {
	r.Logger.Info("publishing message")

	return nil
}

func (r *RabbitMq) Consume(consumer string) error {
	r.Logger.Info("starting consuming")

	return nil
}
