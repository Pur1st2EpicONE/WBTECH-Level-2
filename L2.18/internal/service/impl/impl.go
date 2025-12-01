package impl

import (
	"sort"

	"L2.18/internal/config"
	"L2.18/internal/errs"
	"L2.18/internal/models"
	"L2.18/internal/repository"
	"L2.18/pkg/logger"
)

type Service struct {
	Storage          repository.Storage
	logger           logger.Logger
	maxEventsPerUser int
}

func NewService(config config.Service, storage repository.Storage, logger logger.Logger) *Service {
	return &Service{Storage: storage, logger: logger, maxEventsPerUser: config.MaxEventsPerUser}
}

func (s *Service) CreateEvent(event *models.Event) (string, error) {

	if err := validateCreate(event); err != nil {
		return "", err
	}

	count, err := s.Storage.CountUserEvents(event.Meta.UserID)
	if err != nil {
		return "", err
	}

	if count >= s.maxEventsPerUser {
		return "", errs.ErrMaxEvents
	}

	return s.Storage.CreateEvent(event)

}

func (s *Service) UpdateEvent(event *models.Event) error {

	if err := validateIDs(event.Meta.UserID, event.Meta.EventID); err != nil {
		return err
	}

	if err := validateUpdate(event, s.Storage.GetEventByID(event.Meta.EventID)); err != nil {
		return err
	}

	return s.Storage.UpdateEvent(event)

}

func (s *Service) DeleteEvent(meta *models.Meta) error {

	if err := validateIDs(meta.UserID, meta.EventID); err != nil {
		return err
	}

	if err := validateDelete(meta, s.Storage.GetEventByID(meta.EventID)); err != nil {
		return err
	}

	return s.Storage.DeleteEvent(meta)

}

func (s *Service) GetEvents(meta *models.Meta, period models.Period) ([]models.Event, error) {

	if err := validateGet(meta); err != nil {
		return nil, err
	}

	events, err := s.Storage.GetEvents(meta, period)
	if err != nil {
		return nil, err
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Meta.EventDate.After(events[j].Meta.EventDate)
	})

	return events, nil
}
