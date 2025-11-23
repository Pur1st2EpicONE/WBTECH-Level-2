package memory

import (
	"fmt"
	"time"

	"L2.18/internal/config"
	"L2.18/internal/models"
	"L2.18/pkg/logger"
	"github.com/google/uuid"
)

type Storage struct {
	Hm               map[int]map[string][]*models.Event
	UserEvents       map[int]int
	logger           logger.Logger
	maxEventsPerUser int
}

func NewStorage(config config.Storage, logger logger.Logger) *Storage {
	return &Storage{
		Hm:               make(map[int]map[string][]*models.Event, config.MaxUsers),
		UserEvents:       make(map[int]int),
		logger:           logger,
		maxEventsPerUser: config.MaxEventsPerUser,
	}
}

func (s *Storage) CreateEvent(event *models.Event) error {

	if s.UserEvents[event.UserID] >= s.maxEventsPerUser {
		return fmt.Errorf("max events limit reached")
	}

	if _, userExists := s.Hm[event.UserID]; !userExists {
		s.Hm[event.UserID] = make(map[string][]*models.Event)
	}

	eventDate := event.Date.Format("2006-01-02")

	if event.EventID == "" {
		event.EventID = uuid.New().String()
	}

	s.Hm[event.UserID][eventDate] = append(s.Hm[event.UserID][eventDate], event)
	s.UserEvents[event.UserID]++

	return nil

}

func (s *Storage) GetEventsForDay(userID int, date time.Time) ([]*models.Event, error) {
	userEvents, ok := s.Hm[userID]
	if !ok {
		return nil, nil
	}
	return userEvents[date.Format("2006-01-02")], nil
}

func (s *Storage) GetEventsForWeek(userID int, date time.Time) ([]*models.Event, error) {

}

func (s *Storage) GetEventsForMonth(userID int, date time.Time) ([]*models.Event, error) {

}

func (s *Storage) UpdateEvent(event *models.Event) error {
	eventDate := event.Date.Format("2006-01-02")

	return fmt.Errorf("event not found")
}

func (s *Storage) DeleteEvent(userID int, day time.Time, eventID string) error {
	eventDate := day.Format("2006-01-02")

	return fmt.Errorf("event not found")
}

func (s *Storage) Close() error {

}
