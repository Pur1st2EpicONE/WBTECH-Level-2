package impl

import (
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

func (s *Service) CreateEvent(data *models.Data) (string, error) {
	return s.Storage.CreateEvent(data)
}

func (s *Service) UpdateEvent(data *models.Data) error {
	return s.Storage.UpdateEvent(data)
}

func (s *Service) DeleteEvent(meta *models.Meta) error {
	return s.Storage.DeleteEvent(meta)
}

func (s *Service) GetEventsForDay(meta *models.Meta) ([]models.Event, error) {
	return s.Storage.GetEventsForDay(meta)
}

func (s *Service) GetEventsForWeek(meta *models.Meta) ([]models.Event, error) {
	return s.Storage.GetEventsForWeek(meta)
}

func (s *Service) GetEventsForMonth(meta *models.Meta) ([]models.Event, error) {
	return s.Storage.GetEventsForMonth(meta)
}
