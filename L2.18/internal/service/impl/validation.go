package impl

import (
	"fmt"
	"time"

	"L2.18/internal/errs"
	"L2.18/internal/models"
)

func validateCreate(data *models.Data) error {

	if data.Meta.UserID <= 0 {
		return errs.ErrInvalidUserID
	}

	if err := validateDate(data.Meta.CurrentDate); err != nil {
		return err
	}

	if err := validateEvent(data.Event); err != nil {
		return err
	}

	return nil
}

func validateUpdate(data *models.Data) error {

	if data.Meta.UserID <= 0 {
		return errs.ErrInvalidUserID
	}

	if data.Meta.EventID == "" {
		return errs.ErrMissingEventID
	}

	if data.Meta.CurrentDate.IsZero() {
		return errs.ErrMissingDate
	}

	if data.Meta.NewDate.IsZero() {
		return errs.ErrMissingDate
	}

	if err := validateDate(data.Meta.NewDate); err != nil {
		return err
	}

	if err := validateEvent(data.Event); err != nil {
		return err
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

	if meta.CurrentDate.IsZero() {
		return errs.ErrMissingDate
	}

	return nil

}

func validateGet(meta *models.Meta) error {

	if meta.UserID <= 0 {
		return errs.ErrInvalidUserID
	}

	if meta.CurrentDate.IsZero() {
		return errs.ErrMissingDate
	}

	return nil

}

func validateDate(date time.Time) error {

	if date.IsZero() {
		return errs.ErrMissingDate
	}

	eventUTC := date.UTC().Truncate(24 * time.Hour)
	todayUTC := time.Now().UTC().Truncate(24 * time.Hour)

	if eventUTC.Before(todayUTC) {
		return fmt.Errorf("%w: date %s", errs.ErrEventInPast, eventUTC.Format("2006-01-02"))

	}

	if eventUTC.After(todayUTC.AddDate(10, 0, 0)) {
		return fmt.Errorf("%w: date %s", errs.ErrEventTooFar, eventUTC.Format("2006-01-02"))
	}

	return nil

}

func validateEvent(event models.Event) error {

	if len(event.Text) == 0 {
		return errs.ErrEmptyEventText
	}

	if len(event.Text) > 500 {
		return errs.ErrEventTextTooLong
	}

	return nil

}
