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
	db               map[int]map[string][]*models.Data
	userEventCount   map[int]int
	logger           logger.Logger
	maxEventsPerUser int
}

func NewStorage(config config.Storage, logger logger.Logger) *Storage {
	return &Storage{
		db:               make(map[int]map[string][]*models.Data, config.MaxUsers),
		userEventCount:   make(map[int]int),
		logger:           logger,
		maxEventsPerUser: config.MaxEventsPerUser,
	}
}

func (s *Storage) CreateEvent(d *models.Data) (string, error) {

	if s.userEventCount[d.Meta.UserID] >= s.maxEventsPerUser {
		return "", fmt.Errorf("max events limit reached")
	}

	if _, userExists := s.db[d.Meta.UserID]; !userExists {
		s.db[d.Meta.UserID] = make(map[string][]*models.Data)
	}

	eventDate := format(d.Meta.EventDate)

	if d.Meta.EventID == "" {
		d.Meta.EventID = uuid.New().String()
	}

	s.db[d.Meta.UserID][eventDate] = append(s.db[d.Meta.UserID][eventDate], d)
	s.userEventCount[d.Meta.UserID]++

	return d.Meta.EventID, nil

}

func (s *Storage) UpdateEvent(updated *models.Data) error {

	currDate := format(updated.Meta.EventDate)
	newDate := format(updated.Meta.NewDate)

	allUserEvents, userFound := s.db[updated.Meta.UserID]
	if !userFound {
		return fmt.Errorf("user not found")
	}

	dayEvents, eventsFound := allUserEvents[currDate]
	if !eventsFound {
		return fmt.Errorf("no events found for this date")
	}

	for i, current := range dayEvents {

		if current.Meta.EventID == updated.Meta.EventID {

			if currDate == newDate {
				current.Event = updated.Event
				return nil
			}

			copy(dayEvents[i:], dayEvents[i+1:])
			dayEvents[len(dayEvents)-1] = nil
			dayEvents = dayEvents[:len(dayEvents)-1]

			if len(dayEvents) == 0 {
				delete(allUserEvents, currDate)
			} else {
				allUserEvents[currDate] = dayEvents
			}

			allUserEvents[newDate] = append(allUserEvents[newDate], updated)

			return nil

		}

	}

	return fmt.Errorf("event not found")

}

func (s *Storage) DeleteEvent(m *models.Meta) error {

	date := format(m.EventDate)

	allUserEvents, userFound := s.db[m.UserID]
	if !userFound {
		return fmt.Errorf("user not found")
	}

	dayEvents, eventsFound := allUserEvents[date]
	if !eventsFound {
		return fmt.Errorf("no events for this date found")
	}

	for i, e := range dayEvents {

		if e.Meta.EventID == m.EventID {

			copy(dayEvents[i:], dayEvents[i+1:])
			dayEvents[len(dayEvents)-1] = nil
			dayEvents = dayEvents[:len(dayEvents)-1]

			if len(dayEvents) == 0 {
				delete(allUserEvents, date)
			} else {
				allUserEvents[date] = dayEvents
			}

			s.userEventCount[m.UserID]--

			return nil

		}

	}

	return fmt.Errorf("event not found")

}

func (s *Storage) GetEventsForDay(meta *models.Meta) ([]models.Event, error) {

	allUserEvents, eventsFound := s.db[meta.UserID]
	if !eventsFound {
		return []models.Event{}, nil
	}

	dayEvents, eventsFound := allUserEvents[format(meta.EventDate)]
	if !eventsFound {
		return []models.Event{}, nil
	}

	res := make([]models.Event, len(dayEvents))
	for i, e := range dayEvents {
		res[i] = e.Event
	}

	return res, nil

}

func (s *Storage) GetEventsForWeek(meta *models.Meta) ([]models.Event, error) {

	allUserEvents, eventsFound := s.db[meta.UserID]
	if !eventsFound {
		return []models.Event{}, nil
	}

	var res []models.Event
	year, week := meta.EventDate.ISOWeek()

	for _, dayEvents := range allUserEvents {
		for _, entry := range dayEvents {
			entryYear, entryWeek := entry.Meta.EventDate.ISOWeek()
			if entryYear == year && entryWeek == week {
				res = append(res, entry.Event)
			}
		}
	}

	return res, nil

}

func (s *Storage) GetEventsForMonth(meta *models.Meta) ([]models.Event, error) {

	allUserEvents, eventsFound := s.db[meta.UserID]
	if !eventsFound {
		return []models.Event{}, nil
	}

	var res []models.Event
	month := meta.EventDate.Month()
	year := meta.EventDate.Year()

	for _, dayEvents := range allUserEvents {
		for _, entry := range dayEvents {
			entryMonth := entry.Meta.EventDate.Month()
			entryYear := entry.Meta.EventDate.Year()
			if entryYear == year && entryMonth == month {
				res = append(res, entry.Event)
			}
		}
	}

	return res, nil

}

func (s *Storage) Close() error {
	s.db = nil
	s.userEventCount = nil
	return nil
}

func format(date time.Time) string {
	return date.Format("2006-01-02")
}
