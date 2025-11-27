package service

import (
	"L2.18/internal/models"
	"L2.18/internal/repository"
	"L2.18/internal/service/impl"
	"L2.18/pkg/logger"
)

type Service interface {
	CreateEvent(data *models.Data) (string, error)
	UpdateEvent(data *models.Data) error
	DeleteEvent(meta *models.Meta) error
	GetEventsForDay(meta *models.Meta) ([]models.Event, error)
	GetEventsForWeek(meta *models.Meta) ([]models.Event, error)
	GetEventsForMonth(meta *models.Meta) ([]models.Event, error)
}

func NewService(storage repository.Storage, logger logger.Logger) Service {
	return impl.NewService(storage, logger)
}
