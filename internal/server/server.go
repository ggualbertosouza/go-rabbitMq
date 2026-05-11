package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	Port string
}

func (s *Server) Init(handlers http.Handler) {
	httpServer := &http.Server{
		Addr:         ":" + s.Port,
		Handler:      handlers,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
