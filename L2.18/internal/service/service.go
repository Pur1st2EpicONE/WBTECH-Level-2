package service

import (
	"L2.18/internal/config"
	"L2.18/internal/models"
	"L2.18/internal/repository"
	"L2.18/internal/service/impl"
	"L2.18/pkg/logger"
)

type Service interface {
	CreateEvent(event *models.Event) (string, error)
	UpdateEvent(event *models.Event) error
	DeleteEvent(meta *models.Meta) error
	GetEvents(meta *models.Meta, period models.Period) ([]models.Event, error)
}

func NewService(config config.Service, storage repository.Storage, logger logger.Logger) Service {
	return impl.NewService(config, storage, logger)
}
