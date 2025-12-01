package repository

import (
	"L2.18/internal/config"
	"L2.18/internal/models"
	"L2.18/internal/repository/memory"
	"L2.18/pkg/logger"
)

type Storage interface {
	CreateEvent(event *models.Event) (string, error)
	UpdateEvent(event *models.Event) error
	DeleteEvent(meta *models.Meta) error
	GetEventByID(eventID string) *models.Event
	CountUserEvents(userID int) (int, error)
	GetEvents(meta *models.Meta, period models.Period) ([]models.Event, error)
	Close() error
}

func NewStorage(db any, config config.Storage, logger logger.Logger) Storage {
	if db == nil {
		return memory.NewStorage(config, logger)
	} else {
		panic("unsupported storage type")
	}
}
