package service

import (
	"L2.18/internal/config"
	"L2.18/internal/models"
	"L2.18/internal/repository"
	"L2.18/internal/service/impl"
	"L2.18/pkg/logger"
)

type Service interface {
	CreateEvent(data *models.Event) (string, error)
	UpdateEvent(data *models.Event) error
	DeleteEvent(meta *models.Meta) error
	GetEventsForDay(meta *models.Meta) ([]models.Event, error)
	GetEventsForWeek(meta *models.Meta) ([]models.Event, error)
	GetEventsForMonth(meta *models.Meta) ([]models.Event, error)
}

func NewService(config config.Service, storage repository.Storage, logger logger.Logger) Service {
	return impl.NewService(config, storage, logger)
}
