package memory

import (
	"fmt"
	"sync"
	"time"

	"L2.18/internal/config"
	"L2.18/internal/models"
	"L2.18/pkg/logger"
	"github.com/google/uuid"
)

type Storage struct {
	db             map[int]map[string][]*models.Event
	eventsByID     map[string]*models.Event
	userEventCount map[int]int
	logger         logger.Logger
	mu             sync.RWMutex
}

func NewStorage(config config.Storage, logger logger.Logger) *Storage {
	return &Storage{
		db:             make(map[int]map[string][]*models.Event, config.ExpectedUsers),
		eventsByID:     make(map[string]*models.Event, config.ExpectedUsers),
		userEventCount: make(map[int]int, config.ExpectedUsers),
		logger:         logger,
	}
}

func (s *Storage) CreateEvent(event *models.Event) (string, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, userExists := s.db[event.Meta.UserID]; !userExists {
		s.db[event.Meta.UserID] = make(map[string][]*models.Event)
	}

	eventDate := format(event.Meta.EventDate)
	event.Meta.EventID = uuid.New().String()

	s.db[event.Meta.UserID][eventDate] = append(s.db[event.Meta.UserID][eventDate], event)
	s.userEventCount[event.Meta.UserID]++

	if s.eventsByID == nil {
		s.eventsByID = make(map[string]*models.Event)
	}
	s.eventsByID[event.Meta.EventID] = event

	return event.Meta.EventID, nil

}

func (s *Storage) UpdateEvent(new *models.Event) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	current := s.eventsByID[new.Meta.EventID]

	if current.Data != new.Data {
		updateData(&current.Data, &new.Data)
	}

	if !new.Meta.NewDate.IsZero() && !new.Meta.NewDate.Equal(current.Meta.EventDate) {

		newDate := format(new.Meta.NewDate)
		oldDate := format(current.Meta.EventDate)

		dayEvents := s.db[current.Meta.UserID][oldDate]

		for i, e := range dayEvents {

			if e.Meta.EventID == current.Meta.EventID {
				copy(dayEvents[i:], dayEvents[i+1:])
				dayEvents[len(dayEvents)-1] = nil
				dayEvents = dayEvents[:len(dayEvents)-1]
				break
			}

		}

		if len(dayEvents) == 0 {
			delete(s.db[current.Meta.UserID], oldDate)
		} else {
			s.db[current.Meta.UserID][oldDate] = dayEvents
		}

		current.Meta.EventDate = new.Meta.NewDate
		s.db[current.Meta.UserID][newDate] = append(s.db[current.Meta.UserID][newDate], current)

	}

	return nil

}

func (s *Storage) DeleteEvent(meta *models.Meta) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	current := s.eventsByID[meta.EventID]
	date := format(current.Meta.EventDate)

	userID := current.Meta.UserID
	dayEvents := s.db[userID][date]

	for i, e := range dayEvents {

		if e.Meta.EventID == meta.EventID {
			copy(dayEvents[i:], dayEvents[i+1:])
			dayEvents[len(dayEvents)-1] = nil
			dayEvents = dayEvents[:len(dayEvents)-1]
			break
		}

	}

	if len(dayEvents) == 0 {
		delete(s.db[userID], date)
	} else {
		s.db[userID][date] = dayEvents
	}

	s.userEventCount[userID]--
	delete(s.eventsByID, meta.EventID)

	return nil

}

func (s *Storage) GetEventByID(eventID string) *models.Event {

	s.mu.RLock()
	defer s.mu.RUnlock()

	if event, eventFound := s.eventsByID[eventID]; eventFound {
		return event
	}

	return nil

}

func (s *Storage) CountUserEvents(userID int) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.userEventCount[userID], nil
}

func (s *Storage) GetEvents(meta *models.Meta, period models.Period) ([]models.Event, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	if period != models.Day && period != models.Week && period != models.Month {
		return nil, fmt.Errorf("unknown period: %s", period)
	}

	allUserEvents, eventsFound := s.db[meta.UserID]
	if !eventsFound {
		return []models.Event{}, nil
	}

	switch period {
	case models.Day:
		return s.getEventsForDay(allUserEvents, meta)

	case models.Week:
		return s.getEventsForWeek(allUserEvents, meta)

	case models.Month:
		return s.getEventsForMonth(allUserEvents, meta)
	}

	return []models.Event{}, nil

}

func (s *Storage) getEventsForDay(allUserEvents map[string][]*models.Event, meta *models.Meta) ([]models.Event, error) {

	dayEvents, eventsFound := allUserEvents[format(meta.EventDate)]
	if !eventsFound {
		return []models.Event{}, nil
	}

	res := make([]models.Event, len(dayEvents))
	for i, event := range dayEvents {
		res[i] = *event
	}

	return res, nil

}

func (s *Storage) getEventsForWeek(allUserEvents map[string][]*models.Event, meta *models.Meta) ([]models.Event, error) {

	var res []models.Event
	year, week := meta.EventDate.ISOWeek()

	for _, dayEvents := range allUserEvents {
		for _, entry := range dayEvents {
			entryYear, entryWeek := entry.Meta.EventDate.ISOWeek()
			if entryYear == year && entryWeek == week {
				res = append(res, *entry)
			}
		}
	}

	return res, nil

}

func (s *Storage) getEventsForMonth(allUserEvents map[string][]*models.Event, meta *models.Meta) ([]models.Event, error) {

	month := meta.EventDate.Month()
	year := meta.EventDate.Year()
	var res []models.Event

	for _, dayEvents := range allUserEvents {
		for _, entry := range dayEvents {
			entryMonth := entry.Meta.EventDate.Month()
			entryYear := entry.Meta.EventDate.Year()
			if entryYear == year && entryMonth == month {
				res = append(res, *entry)
			}
		}
	}

	return res, nil

}

func (s *Storage) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.db = nil
	s.eventsByID = nil
	s.userEventCount = nil
	s.logger.LogInfo("in-memory storage — cleared and stopped", "layer", "repository.memory")
	return nil
}

func updateData(current *models.Data, new *models.Data) {
	current.Text = new.Text
}

func format(date time.Time) string {
	return date.Format("2006-01-02")
}
