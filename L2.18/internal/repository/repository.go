package repository

import (
	"time"

	"L2.18/internal/config"
	"L2.18/internal/models"
	"L2.18/internal/repository/memory"
	"L2.18/pkg/logger"
)

type Storage interface {
	CreateEvent(e *models.Event) error
	UpdateEvent(e *models.Event) error
	DeleteEvent(int, time.Time, string) error
	GetEventsForDay(userID int, date time.Time) ([]*models.Event, error)
	GetEventsForWeek(userID int, date time.Time) ([]*models.Event, error)
	GetEventsForMonth(userID int, date time.Time) ([]*models.Event, error)
	Close() error
}

func NewStorage(config config.Storage, logger logger.Logger) Storage {
	return memory.NewStorage(config, logger)
}
