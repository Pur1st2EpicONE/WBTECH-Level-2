package httpserver

import (
	"context"
	"net/http"
	"time"

	"L2.18/internal/config"
	"L2.18/pkg/logger"
)

type HttpServer struct {
	Srv             *http.Server
	ShutdownTimeout time.Duration
}

func NewServer(config config.Server, handler http.Handler) *HttpServer {
	server := new(HttpServer)
	server.Srv = &http.Server{
		Addr:           ":" + config.Port,
		Handler:        handler,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}
	server.ShutdownTimeout = config.ShutdownTimeout
	return server
}

func (s *HttpServer) Run(ctx context.Context, logger logger.Logger) error {
	logger.LogInfo("server — receiving requests", "layer", "server")
	return s.Srv.ListenAndServe()
}

func (s *HttpServer) Shutdown(ctx context.Context, logger logger.Logger) {
	if err := s.Srv.Shutdown(ctx); err != nil {
		logger.LogError("server — failed to shutdown gracefully", err, "layer", "server")
	} else {
		logger.LogInfo("server — shutdown complete", "layer", "server")
	}
}
