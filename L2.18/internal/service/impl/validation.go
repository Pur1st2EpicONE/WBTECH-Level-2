package impl

import (
	"fmt"
	"time"

	"L2.18/internal/errs"
	"L2.18/internal/models"
	"github.com/google/uuid"
)

func validateCreate(data *models.Event) error {

	if data.Meta.UserID <= 0 {
		return errs.ErrInvalidUserID
	}

	if err := validateDate(data.Meta.EventDate); err != nil {
		return err
	}

	if err := validateEvent(data.Data); err != nil {
		return err
	}

	return nil

}

func validateUpdate(data *models.Event) error {

	if data.Meta.UserID <= 0 {
		return errs.ErrInvalidUserID
	}

	if err := validateEventID(data.Meta.EventID); err != nil {
		return err
	}

	if data.Meta.NewDate.IsZero() && data.Data.Text == "" {
		return errs.ErrNothingToUpdate
	}

	if !data.Meta.NewDate.IsZero() {
		if err := validateDate(data.Meta.NewDate); err != nil {
			return err
		}
	}

	if data.Data.Text != "" {
		if err := validateEvent(data.Data); err != nil {
			return err
		}
	}

	return nil
}

func validateDelete(meta *models.Meta) error {

	if meta.UserID <= 0 {
		return errs.ErrInvalidUserID
	}

	if meta.EventID == "" {
		return errs.ErrMissingEventID
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

	if date.IsZero() { // ?????
		return errs.ErrMissingDate // ?????
	} // ?????

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

func validateEvent(event models.Data) error {

	if len(event.Text) == 0 {
		return errs.ErrEmptyEventText
	}

	if len(event.Text) > 500 {
		return errs.ErrEventTextTooLong
	}

	return nil

}

func validateEventID(id string) error {

	if id == "" {
		return errs.ErrMissingEventID
	}
	_, err := uuid.Parse(id)
	if err != nil {
		return errs.ErrInvalidEventID
	}

	return nil

}
