package httpserver

import (
	"context"
	"net/http"
	"time"

	"L2.18/internal/config"
	"L2.18/pkg/logger"
)

type HttpServer struct {
	srv             *http.Server
	shutdownTimeout time.Duration
	logger          logger.Logger
}

func NewServer(config config.Server, handler http.Handler, logger logger.Logger) *HttpServer {
	server := new(HttpServer)
	server.srv = &http.Server{
		Addr:           ":" + config.Port,
		Handler:        handler,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}
	server.shutdownTimeout = config.ShutdownTimeout
	server.logger = logger
	return server
}

func (s *HttpServer) Run() error {
	s.logger.LogInfo("server — receiving requests", "layer", "server")
	return s.srv.ListenAndServe()
}

func (s *HttpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.LogError("server — failed to shutdown gracefully", err, "layer", "server")
	} else {
		s.logger.LogInfo("server — shutdown complete", "layer", "server")
	}
}
