package main

import (
	config "github.com/ggualbertosouza/go-rabbitMq/internal/config/env"
	"github.com/ggualbertosouza/go-rabbitMq/internal/config/logger"
	"github.com/ggualbertosouza/go-rabbitMq/internal/rabbitmq"
	"github.com/ggualbertosouza/go-rabbitMq/internal/server"
)

func main() {
	log := logger.NewLogger()

	config.LoadEnv()

	initRabbit(log)
	initServer(log)
}

func initServer(log logger.Logger) {
	hc := config.HttpConfig()

	s := server.Server{
		Port:   hc.Port,
		Logger: log,
	}
	r := server.NewRouter()
	s.Init(r)
}

func initRabbit(log logger.Logger) {
	rc := config.RabbitMQConfig()
	rabbit, err := rabbitmq.New(log, rc.URL)
	if err != nil {
		panic(err)
	}

	err = rabbit.Init(rc.Exchange, rc.Queue, rc.RoutingKey)
	if err != nil {
		panic(err)
	}
}
