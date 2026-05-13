package main

import (
	config "github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/env"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/rabbitmq"
	"github.com/ggualbertosouza/go-rabbitMq/internal/presentation/workers"
)

func main() {
	log := logger.NewLogger("Consumers")
	log.Info("initializing consumers")

	config.LoadEnv()

	rabbit, err := initRabbit(log)
	if err != nil {
		panic(err)
	}

	wk := workers.New(log, rabbit)
	wk.Init()
}

func initRabbit(log logger.Logger) (*rabbitmq.RabbitMq, error) {
	rc := config.Rabbitmq()
	rabbit, err := rabbitmq.New(log, rc.URL)
	if err != nil {
		return nil, err
	}

	return rabbit, nil
}
