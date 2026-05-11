package main

import (
	"log"

	"github.com/ggualbertosouza/go-rabbitMq/internal/config"
	"github.com/ggualbertosouza/go-rabbitMq/internal/rabbitmq"
	"github.com/ggualbertosouza/go-rabbitMq/internal/server"
)

func main() {
	config.LoadEnv()

	hc := config.HttpConfig()
	rc := config.RabbitMQConfig()

	rabbit := rabbitmq.New(rc.URL)
	if err := rabbit.Setup(rc.Exchange, rc.Queue, rc.RoutingKey); err != nil {
		log.Fatal(err)
	}

	s := server.Server{Port: hc.Port}
	r := server.NewRouter()
	s.Init(r)
}
