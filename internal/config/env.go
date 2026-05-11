package config

import (
	"os"

	"github.com/joho/godotenv"
)

type HttpServer struct {
	Port string
}

func HttpConfig() *HttpServer {
	return &HttpServer{
		Port: get("SERVER_PORT", "8080"),
	}
}

func LoadEnv(files ...string) error {
	if len(files) == 0 {
		return godotenv.Load()
	}

	return godotenv.Load(files...)
}

func get(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
