package main

import (
	config "github.com/ggualbertosouza/go-rabbitMq/internal/config/env"
	"github.com/ggualbertosouza/go-rabbitMq/internal/config/logger"
	"github.com/ggualbertosouza/go-rabbitMq/internal/rabbitmq"
	"github.com/ggualbertosouza/go-rabbitMq/internal/server"
	"github.com/ggualbertosouza/go-rabbitMq/internal/server/context"
)

func main() {
	log := logger.NewLogger()

	config.LoadEnv()

	rabbit, err := initRabbit(log)
	if err != nil {
		panic(err)
	}

	deps := context.Dependencies{
		Logger: log,
		Rabbit: rabbit,
	}

	initRabbitConsumer(log, rabbit)
	initServer(log, deps)
}

func initServer(
	log logger.Logger,
	deps context.Dependencies,
) {
	hc := config.HttpConfig()

	s := server.Server{
		Port:   hc.Port,
		Logger: log,
	}
	r := server.NewRouter(deps)
	s.Init(r)
}

func initRabbit(log logger.Logger) (*rabbitmq.RabbitMq, error) {
	rc := config.RabbitMQConfig()
	rabbit, err := rabbitmq.New(log, rc.URL)
	if err != nil {
		return nil, err
	}

	return rabbit, nil
}

func initRabbitConsumer(
	log logger.Logger,
	rabbit *rabbitmq.RabbitMq,
) {
	err := rabbit.Consume(string(rabbitmq.UserQueue1), "userConsumerQueue1")
	if err != nil {
		log.Error(err.Error())
	}
}
