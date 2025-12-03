package impl

import (
	"fmt"
	"time"

	"L2.18/internal/errs"
	"L2.18/internal/models"
	"github.com/google/uuid"
)

func validateCreate(event *models.Event) error {

	if event.Meta.UserID <= 0 {
		return errs.ErrInvalidUserID
	}

	if err := validateDate(event.Meta.EventDate); err != nil {
		return err
	}

	if err := validateData(event.Data); err != nil {
		return err
	}

	return nil

}

func validateUpdate(event *models.Event, oldEvent *models.Event) error {

	if oldEvent == nil {
		return errs.ErrEventNotFound
	}

	if event.Meta.UserID != oldEvent.Meta.UserID {
		return errs.ErrUnauthorized
	}

	if isNothingToUpdate(event, oldEvent) {
		return errs.ErrNothingToUpdate
	}

	if !event.Meta.NewDate.IsZero() {
		if err := validateDate(event.Meta.NewDate); err != nil {
			return err
		}
	}

	if event.Data.Text != "" {
		if err := validateData(event.Data); err != nil {
			return err
		}
	}

	return nil
}

func isNothingToUpdate(event, oldEvent *models.Event) bool {
	if !event.Meta.NewDate.IsZero() && !oldEvent.Meta.EventDate.Equal(event.Meta.NewDate) {
		return false
	}
	if event.Data.Text != oldEvent.Data.Text {
		return false
	}
	return true
}

func validateDelete(meta *models.Meta, oldEvent *models.Event) error {

	if oldEvent == nil {
		return errs.ErrEventNotFound
	}

	if meta.UserID != oldEvent.Meta.UserID {
		return errs.ErrUnauthorized
	}

	return nil

}

func validateGet(meta *models.Meta) error {

	if meta.UserID <= 0 {
		return errs.ErrInvalidUserID
	}

	if meta.EventDate.IsZero() {
		return errs.ErrMissingDate
	}

	return nil

}

func validateDate(date time.Time) error {

	eventUTC := date.UTC().Truncate(24 * time.Hour)
	todayUTC := time.Now().UTC().Truncate(24 * time.Hour)

	if eventUTC.Before(todayUTC) {
		return fmt.Errorf("%w: %s", errs.ErrEventInPast, eventUTC.Format("2006-01-02"))

	}

	if eventUTC.After(todayUTC.AddDate(10, 0, 0)) {
		return fmt.Errorf("%w: %s", errs.ErrEventTooFar, eventUTC.Format("2006-01-02"))
	}

	return nil

}

func validateData(data models.Data) error {

	if len(data.Text) > 500 {
		return errs.ErrEventTextTooLong
	}

	return nil

}

func validateIDs(userID int, eventID string) error {

	if userID <= 0 {
		return errs.ErrInvalidUserID
	}

	if eventID == "" {
		return errs.ErrMissingEventID
	}

	_, err := uuid.Parse(eventID)
	if err != nil {
		return errs.ErrInvalidEventID
	}

	return nil

}
