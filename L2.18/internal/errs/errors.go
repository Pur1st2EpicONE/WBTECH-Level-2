package errs

import "errors"

var (
	ErrInvalidJSON       = errors.New("invalid JSON")                                     // invalid JSON
	ErrMissingParams     = errors.New("missing required params")                          // missing required params
	ErrInvalidUserID     = errors.New("invalid user id")                                  // invalid user id
	ErrInvalidDateFormat = errors.New("invalid date format")                              // invalid date format
	ErrInvalidQuery      = errors.New("invalid query parameters")                         // invalid query parameters
	ErrEmptyEventText    = errors.New("event text cannot be empty")                       // event text cannot be empty
	ErrEventTextTooLong  = errors.New("event text exceeds maximum length")                // event text exceeds maximum length
	ErrEventInPast       = errors.New("event cannot be in the past")                      // event cannot be in the past
	ErrEventTooFar       = errors.New("event cannot be more than 10 years in the future") // event cannot be more than 10 years in the future
	ErrMissingEventID    = errors.New("event id is required")                             // event id is required
	ErrMissingDate       = errors.New("event date is required")                           // event date is required
)

var (
	ErrMaxEvents     = errors.New("max events limit reached") // max events limit reached
	ErrUserNotFound  = errors.New("user not found")           // user not found
	ErrEventNotFound = errors.New("event not found")          // event not found
)
