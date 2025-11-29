package v1

import (
	"time"

	"L2.18/internal/errs"
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
		respondError(c, errs.ErrInvalidJSON)
		return
	}

	eventDate, err := parseDate(request.EventDate)
	if err != nil {
		respondError(c, err)
		return
	}

	event := models.Event{
		Meta: models.Meta{UserID: request.UserID, EventDate: eventDate},
		Data: models.Data{Text: request.Text},
	}

	eventID, err := h.service.CreateEvent(&event)
	if err != nil {
		respondError(c, err)
		return
	}

	respondOK(c, CreateResponseV1{EventID: eventID})

}

func (h *Handler) UpdateEvent(c *gin.Context) {

	var request UpdateRequestV1

	if err := c.ShouldBindJSON(&request); err != nil {
		respondError(c, errs.ErrInvalidJSON)
		return
	}

	var date time.Time
	var err error

	if request.NewDate != nil && *request.NewDate != "" {
		date, err = parseDate(*request.NewDate)
		if err != nil {
			respondError(c, errs.ErrInvalidDateFormat)
			return
		}

	}

	event := models.Event{Meta: models.Meta{UserID: request.UserID, EventID: request.EventID, NewDate: date}}

	if request.Text != nil {
		event.Data.Text = *request.Text
	}

	if err := h.service.UpdateEvent(&event); err != nil {
		respondError(c, err)
		return
	}

	respondOK(c, UpdateResponseV1{Updated: true})

}

func (h *Handler) DeleteEvent(c *gin.Context) {

	var request DeleteRequestV1

	if err := c.ShouldBindJSON(&request); err != nil {
		respondError(c, errs.ErrInvalidJSON)
		return
	}

	meta := models.Meta{UserID: request.UserID, EventID: request.EventID}

	if err := h.service.DeleteEvent(&meta); err != nil {
		respondError(c, err)
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

	userId, eventDate, err := validateQuery(c.Query("user_id"), c.Query("date"))
	if err != nil {
		respondError(c, err)
		return
	}

	events, err := getFunc(&models.Meta{UserID: userId, EventDate: eventDate})
	if err != nil {
		respondError(c, err)
		return
	}

	respEvents := make([]EventDtoV1, len(events))

	for i, e := range events {
		respEvents[i] = EventDtoV1{Text: e.Data.Text, EventDate: e.Meta.EventDate.Format("2006-01-02"), EventID: e.Meta.EventID}
	}

	respondOK(c, ListOfEventsResponseV1{Events: respEvents})

}
