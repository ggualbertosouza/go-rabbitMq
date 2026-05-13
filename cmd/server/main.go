package main

import (
	config "github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/env"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/rabbitmq"
	"github.com/ggualbertosouza/go-rabbitMq/internal/presentation/server"
	serverContext "github.com/ggualbertosouza/go-rabbitMq/internal/presentation/server/context"
)

func main() {
	log := logger.NewLogger("Server")
	log.Info("initializing server")

	config.LoadEnv()

	rabbit, err := initRabbit(log)
	if err != nil {
		panic(err)
	}

	deps := serverContext.Dependencies{
		Logger: log,
		Rabbit: rabbit,
	}

	initServer(log, deps)
}

func initServer(
	log logger.Logger,
	deps serverContext.Dependencies,
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
	rc := config.Rabbitmq()
	rabbit, err := rabbitmq.New(log, rc.URL)
	if err != nil {
		return nil, err
	}

	return rabbit, nil
}
