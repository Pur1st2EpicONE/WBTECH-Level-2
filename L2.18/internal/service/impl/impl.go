package impl

import (
	"time"

	"L2.18/internal/models"
	"L2.18/internal/repository"
	"L2.18/pkg/logger"
)

type Service struct {
	Storage repository.Storage
	logger  logger.Logger
}

func NewService(storage repository.Storage, logger logger.Logger) *Service {
	return &Service{Storage: storage, logger: logger}
}

func (s *Service) CreateEvent(userID int, date time.Time, text string) error {

}

func (s *Service) GetEventsForDay(userID int, date time.Time) ([]models.Event, error) {

}

func (s *Service) GetEventsForWeek(userID int, date time.Time) ([]models.Event, error) {

}

func (s *Service) GetEventsForMonth(userID int, date time.Time) ([]models.Event, error) {

}

func (s *Service) UpdateEvent(e *models.Event) error {

}

func (s *Service) DeleteEvent(eventID int) error {

}
