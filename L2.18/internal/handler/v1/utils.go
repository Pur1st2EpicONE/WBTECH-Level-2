package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"L2.18/internal/errs"
	"github.com/gin-gonic/gin"
)

func validateQuery(userID string, eventDate string) (int, []time.Time, error) {

	if userID == "" || eventDate == "" {
		return 0, nil, fmt.Errorf("%w: user id or date is empty", errs.ErrMissingParams)
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		return 0, nil, fmt.Errorf("%w: %v", errs.ErrInvalidUserID, err)
	}

	dates, err := validateDates(eventDate)
	if err != nil {
		return 0, nil, fmt.Errorf("%w: %v", errs.ErrInvalidDateFormat, err)
	}

	return id, dates, nil

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

func respondOK(c *gin.Context, response any) {
	c.JSON(http.StatusOK, gin.H{"result": response})
}

func respondError(c *gin.Context, err error) {
	if err != nil {
		status, msg := mapErrorToStatus(err)
		c.AbortWithStatusJSON(status, gin.H{"error": msg})
	}
}

func mapErrorToStatus(err error) (int, string) {

	switch {

	case errors.Is(err, errs.ErrInvalidJSON),
		errors.Is(err, errs.ErrInvalidUserID),
		errors.Is(err, errs.ErrInvalidDateFormat),
		errors.Is(err, errs.ErrEmptyEventText),
		errors.Is(err, errs.ErrEventTextTooLong),
		errors.Is(err, errs.ErrEventInPast),
		errors.Is(err, errs.ErrEventTooFar),
		errors.Is(err, errs.ErrMissingEventID),
		errors.Is(err, errs.ErrMissingDate):
		return http.StatusBadRequest, err.Error()

	case errors.Is(err, errs.ErrMaxEvents),
		errors.Is(err, errs.ErrEventNotFound),
		errors.Is(err, errs.ErrUserNotFound):
		return http.StatusServiceUnavailable, err.Error()

	default:
		return http.StatusInternalServerError, "internal server error"
	}

}
