package main

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/config"
	"github.com/ggualbertosouza/go-rabbitMq/internal/server"
)

func main() {
	config.LoadEnv()

	hc := config.HttpConfig()

	s := server.Server{
		Port: hc.Port,
	}
	r := server.NewRouter()

	s.Init(r)
}
