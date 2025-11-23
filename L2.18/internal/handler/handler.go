package handler

import (
	"net/http"

	"L2.18/internal/service"
	"L2.18/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service      service.Service
	logger       logger.Logger
	TemplatePath string
}

func NewHandler(service service.Service, logger logger.Logger) *Handler {
	return &Handler{
		service:      service,
		logger:       logger,
		TemplatePath: "web/templates/*",
	}
}

func (h *Handler) InitRoutes() http.Handler {
	router := gin.New()
	return router
}
