package rabbitmq

type ExchangeType string

const (
	DirectExchange  ExchangeType = "direct"
	TopicExchange   ExchangeType = "topic"
	FanOutExchange  ExchangeType = "fanout"
	HeadersExchange ExchangeType = "headers"
)

const (
	EventMessage string = "event.message"
)
