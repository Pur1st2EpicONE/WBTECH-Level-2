package memory

import (
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
}

func NewStorage(config config.Storage, logger logger.Logger) *Storage {
	return &Storage{
		db:             make(map[int]map[string][]*models.Event, config.MaxUsers),
		eventsByID:     make(map[string]*models.Event, (config.MaxEventsPerUser * config.MaxUsers)),
		userEventCount: make(map[int]int, config.MaxUsers),
		logger:         logger,
	}
}

func (s *Storage) CreateEvent(event *models.Event) (string, error) {

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

	current := s.eventsByID[new.Meta.EventID]

	if current.Data != new.Data {
		updateData(current, new)
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

func (s *Storage) GetMetaByID(eventID string) *models.Meta {
	if data, eventFound := s.eventsByID[eventID]; eventFound {
		return &data.Meta
	}
	return nil
}

func (s *Storage) CountUserEvents(userID int) (int, error) {
	return s.userEventCount[userID], nil
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
	for i, event := range dayEvents {
		res[i] = *event
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
				res = append(res, *entry)
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
				res = append(res, *entry)
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

func updateData(current *models.Event, new *models.Event) {
	if (new.Data.Text != "") && (new.Data.Text != current.Data.Text) {
		current.Data.Text = new.Data.Text
	}
}

func format(date time.Time) string {
	return date.Format("2006-01-02")
}
