package service

import (
	"time"

	"L2.18/internal/models"
	"L2.18/internal/repository"
	"L2.18/internal/service/impl"
	"L2.18/pkg/logger"
)

type Service interface {
	CreateEvent(userID int, date time.Time, text string) error
	UpdateEvent(e *models.Event) error
	DeleteEvent(eventID int) error
	GetEventsForDay(userID int, date time.Time) ([]models.Event, error)
	GetEventsForWeek(userID int, date time.Time) ([]models.Event, error)
	GetEventsForMonth(userID int, date time.Time) ([]models.Event, error)
}

func NewService(storage repository.Storage, logger logger.Logger) Service {
	return impl.NewService(storage, logger)
}
