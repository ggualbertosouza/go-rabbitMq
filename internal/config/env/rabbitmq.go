package config

type RabbitMQ struct {
	URL        string
	Exchange   string
	Queue      string
	RoutingKey string
}

func RabbitMQConfig() *RabbitMQ {
	return &RabbitMQ{
		URL:        get("RABBITMQ_URL", "amqp://admin:admin@localhost:5672/"),
		Exchange:   get("RABBITMQ_EXCHANGE", "events"),
		Queue:      get("RABBITMQ_QUEUE", "messages"),
		RoutingKey: get("RABBITMQ_ROUTING_KEY", "demo.message"),
	}
}
