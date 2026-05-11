package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/ggualbertosouza/go-rabbitMq/internal/config/logger"
)

type Server struct {
	Port   string
	Logger logger.Logger
}

func (s *Server) Init(handlers http.Handler) error {
	httpServer := &http.Server{
		Addr:         ":" + s.Port,
		Handler:      handlers,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.Logger.Info(
		"http server starting",
		logger.String("port", s.Port),
	)

	err := httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.Logger.Error(
			"http server failed",
			logger.Any("error", err),
			logger.String("port", s.Port),
		)

		return err
	}

	s.Logger.Info("http server stopped")

	return nil
}
