package repository

import (
	"L2.18/internal/config"
	"L2.18/internal/models"
	"L2.18/internal/repository/memory"
	"L2.18/pkg/logger"
)

type Storage interface {
	CreateEvent(data *models.Event) (string, error)
	UpdateEvent(data *models.Event) error
	DeleteEvent(meta *models.Meta) error
	GetMetaByID(eventID string) *models.Meta
	CountUserEvents(userID int) (int, error)
	GetEventsForDay(meta *models.Meta) ([]models.Event, error)
	GetEventsForWeek(meta *models.Meta) ([]models.Event, error)
	GetEventsForMonth(meta *models.Meta) ([]models.Event, error)
	Close() error
}

func NewStorage(db any, config config.Storage, logger logger.Logger) Storage {
	if db == nil {
		return memory.NewStorage(config, logger)
	} else {
		panic("unsupported storage type")
	}
}
