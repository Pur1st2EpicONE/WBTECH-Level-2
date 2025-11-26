package server

import (
	"net/http"

	"L2.18/internal/config"
	"L2.18/internal/server/httpserver"
	"L2.18/pkg/logger"
)

type Server interface {
	Run() error
	Shutdown()
}

func NewServer(config config.Server, handler http.Handler, logger logger.Logger) Server {
	return httpserver.NewServer(config, handler, logger)
}
