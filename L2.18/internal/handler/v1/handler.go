package v1

import (
	"fmt"
	"time"

	"L2.18/internal/models"
	"L2.18/internal/service"
	"L2.18/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
	logger  logger.Logger
}

func NewHandler(service service.Service, logger logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateEvent(c *gin.Context) {

	var request CreateRequestV1

	if err := c.ShouldBindJSON(&request); err != nil {
		respondBadRequest(c, err)
		return
	}

	eventDates, err := validateDates(request.EventDate)
	if err != nil {
		respondBadRequest(c, fmt.Errorf("invalid date format, expected YYYY-MM-DD"))
		return
	}

	data := models.Data{
		Meta:  models.Meta{UserID: request.UserID, NewDate: eventDates[0]},
		Event: models.Event{Text: request.Text},
	}

	eventID, err := h.service.CreateEvent(&data)
	if err != nil {
		respondInternalError(c, err)
		return
	}

	respondOK(c, CreateResponseV1{EventID: eventID})

}

func (h *Handler) UpdateEvent(c *gin.Context) {

	var request UpdateRequestV1

	if err := c.ShouldBindJSON(&request); err != nil {
		respondBadRequest(c, err)
		return
	}

	eventDates, err := validateDates(request.CurrentDate, request.NewDate)
	if err != nil {
		respondBadRequest(c, fmt.Errorf("invalid date format, expected YYYY-MM-DD"))
		return
	}

	data := models.Data{
		Meta:  models.Meta{UserID: request.UserID, EventID: request.EventID, CurrentDate: eventDates[0], NewDate: eventDates[1]},
		Event: models.Event{Text: request.Text},
	}

	if err := h.service.UpdateEvent(&data); err != nil {
		respondInternalError(c, err)
		return
	}

	respondOK(c, UpdateResponseV1{Updated: true})

}

func (h *Handler) DeleteEvent(c *gin.Context) {

	var request DeleteRequestV1

	if err := c.ShouldBindJSON(&request); err != nil {
		respondBadRequest(c, err)
		return
	}

	eventDates, err := validateDates(request.EventDate)
	if err != nil {
		respondBadRequest(c, fmt.Errorf("invalid date format, expected YYYY-MM-DD"))
		return
	}

	meta := models.Meta{UserID: request.UserID, CurrentDate: eventDates[0]}

	if err := h.service.DeleteEvent(&meta); err != nil {
		respondInternalError(c, err)
		return
	}

	respondOK(c, DeleteResponseV1{Deleted: true})

}

func (h *Handler) GetEventsForDay(c *gin.Context) {
	h.getEvents(c, h.service.GetEventsForDay)
}

func (h *Handler) GetEventsForWeek(c *gin.Context) {
	h.getEvents(c, h.service.GetEventsForWeek)
}

func (h *Handler) GetEventsForMonth(c *gin.Context) {
	h.getEvents(c, h.service.GetEventsForMonth)
}

func (h *Handler) getEvents(c *gin.Context, getFunc func(*models.Meta) ([]models.Event, error)) {

	var request EventDtoV1

	if err := c.ShouldBindJSON(&request); err != nil {
		respondBadRequest(c, err)
		return
	}

	eventDates, err := validateDates(request.EventDate)
	if err != nil {
		respondBadRequest(c, fmt.Errorf("invalid date format, expected YYYY-MM-DD"))
		return
	}

	meta := models.Meta{UserID: request.UserID, CurrentDate: eventDates[0]}

	events, err := getFunc(&meta)
	if err != nil {
		respondInternalError(c, err)
		return
	}

	respEvents := make([]EventDtoV1, len(events))

	for i, event := range events {
		respEvents[i] = EventDtoV1{
			UserID:    request.UserID,
			EventDate: request.EventDate,
			Text:      event.Text,
		}
	}

	respondOK(c, ListOfEventsResponseV1{Events: respEvents})

}

func validateDates(dates ...string) ([]time.Time, error) {
	res := make([]time.Time, len(dates))
	for i, date := range dates {
		eventDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}
		res[i] = eventDate
	}
	return res, nil
}
