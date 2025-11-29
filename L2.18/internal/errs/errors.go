package errs

import "errors"

var (
	ErrInvalidJSON       = errors.New("invalid JSON")                                     // invalid JSON
	ErrEmptyEventText    = errors.New("event text is required")                           // event text cannot be empty
	ErrMissingDate       = errors.New("event date is required")                           // event date is required
	ErrInvalidDateFormat = errors.New("invalid date format")                              // invalid date format
	ErrEventTextTooLong  = errors.New("event text exceeds maximum length")                // event text exceeds maximum length
	ErrInvalidUserID     = errors.New("missing or invalid user id")                       // missing or invalid user id
	ErrEventInPast       = errors.New("event cannot be in the past")                      // event cannot be in the past: 2025-11-01"
	ErrEventTooFar       = errors.New("event cannot be more than 10 years in the future") // event cannot be more than 10 years in the future
	ErrMaxEvents         = errors.New("max events limit reached")                         // max events limit reached
	ErrNothingToUpdate   = errors.New("nothing to update")                                // nothing to update
	ErrEventNotFound     = errors.New("event not found")                                  // event not found
	ErrInvalidEventID    = errors.New("invalid event_id")                                 // invalid event_id
	ErrUnauthorized      = errors.New("you have no right to change this event")           // you have no right to change this event
)

var (
	ErrUserNotFound   = errors.New("user not found")           // user not found
	ErrMissingParams  = errors.New("missing required params")  // missing required params
	ErrInvalidQuery   = errors.New("invalid query parameters") // invalid query parameters
	ErrMissingEventID = errors.New("event id is required")     // event id is required
	ErrInternal       = errors.New("internal server error")    // internal server error
)
