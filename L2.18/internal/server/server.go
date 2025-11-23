package server

import (
	"context"
	"net/http"

	"L2.18/internal/config"
	"L2.18/internal/server/httpserver"
	"L2.18/pkg/logger"
)

type Server interface {
	Run(ctx context.Context, logger logger.Logger) error
	Shutdown(ctx context.Context, logger logger.Logger)
}

func NewServer(cfg config.Server, handler http.Handler) Server {
	return httpserver.NewServer(cfg, handler)
}
