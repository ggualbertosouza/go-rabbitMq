package config

type HttpServer struct {
	Port string
}

func HttpConfig() *HttpServer {
	return &HttpServer{
		Port: get("SERVER_PORT", "8080"),
	}
}
