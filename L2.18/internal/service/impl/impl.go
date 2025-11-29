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

func (s *Service) UpdateEvent(data *models.Event) error {

	if err := validateUpdate(data); err != nil {
		return err
	}

	oldMeta := s.Storage.GetMetaByID(data.Meta.EventID)
	if oldMeta == nil {
		return errs.ErrEventNotFound
	}

	if oldMeta.UserID != data.Meta.UserID {
		return errs.ErrUnauthorized
	}

	return s.Storage.UpdateEvent(data)

}

func (s *Service) DeleteEvent(meta *models.Meta) error {

	if err := validateDelete(meta); err != nil {
		return err
	}

	oldMeta := s.Storage.GetMetaByID(meta.EventID)
	if oldMeta == nil {
		return errs.ErrEventNotFound
	}

	if oldMeta.UserID != meta.UserID {
		return errs.ErrUnauthorized
	}

	return s.Storage.DeleteEvent(meta)

}

func (s *Service) GetEventsForDay(meta *models.Meta) ([]models.Event, error) {

	if err := validateGet(meta); err != nil {
		return nil, err
	}

	events, err := s.Storage.GetEventsForDay(meta)
	if err != nil {
		return nil, err
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Meta.EventDate.After(events[j].Meta.EventDate)
	})

	return events, nil

}

func (s *Service) GetEventsForWeek(meta *models.Meta) ([]models.Event, error) {

	if err := validateGet(meta); err != nil {
		return nil, err
	}

	events, err := s.Storage.GetEventsForWeek(meta)
	if err != nil {
		return nil, err
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Meta.EventDate.After(events[j].Meta.EventDate)
	})

	return events, nil

}

func (s *Service) GetEventsForMonth(meta *models.Meta) ([]models.Event, error) {

	if err := validateGet(meta); err != nil {
		return nil, err
	}

	events, err := s.Storage.GetEventsForMonth(meta)
	if err != nil {
		return nil, err
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Meta.EventDate.After(events[j].Meta.EventDate)
	})

	return events, nil
}
