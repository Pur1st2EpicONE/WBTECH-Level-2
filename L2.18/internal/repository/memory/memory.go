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

func (s *Storage) UpdateEvent(d *models.Data) error {

	currDate := format(d.Meta.EventDate)
	newDate := format(d.Meta.UpdateDate)

	allUserEvents, ok := s.db[d.Meta.UserID]
	if !ok {
		return fmt.Errorf("user not found")
	}

	dateEvents, ok := allUserEvents[currDate]
	if !ok {
		return fmt.Errorf("no events for this date found")
	}

	for i, e := range dateEvents {

		if e.Meta.EventID == d.Meta.EventID {

			if currDate == newDate {
				e.Event = d.Event
				return nil
			}

			copy(dateEvents[i:], dateEvents[i+1:])
			dateEvents[len(dateEvents)-1] = nil
			dateEvents = dateEvents[:len(dateEvents)-1]

			if len(dateEvents) == 0 {
				delete(allUserEvents, currDate)
			} else {
				allUserEvents[currDate] = dateEvents
			}

			allUserEvents[newDate] = append(allUserEvents[newDate], d)

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

	dateEvents, eventsFound := allUserEvents[date]
	if !eventsFound {
		return fmt.Errorf("no events for this date found")
	}

	for i, e := range dateEvents {

		if e.Meta.EventID == m.EventID {

			copy(dateEvents[i:], dateEvents[i+1:])
			dateEvents[len(dateEvents)-1] = nil
			dateEvents = dateEvents[:len(dateEvents)-1]

			if len(dateEvents) == 0 {
				delete(allUserEvents, date)
			} else {
				allUserEvents[date] = dateEvents
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
