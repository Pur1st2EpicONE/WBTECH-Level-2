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
		respondBadRequest(c, err) // !
		return
	}

	eventDate, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		respondBadRequest(c, fmt.Errorf("invalid date format, expected YYYY-MM-DD"))
		return
	}

	data := models.Data{
		Meta:  models.Meta{UserID: request.UserID, EventDate: eventDate},
		Event: models.Event{Text: request.Text},
	}

	eventID, err := h.service.CreateEvent(&data)
	if err != nil {
		respondInternalError(c, err) // || businessError?
		return
	}

	respondOK(c, eventID)

}
