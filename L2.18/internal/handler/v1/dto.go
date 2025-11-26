package v1

import "time"

type CreateRequestV1 struct {
	UserID int    `json:"user_id" binding:"required"`
	Text   string `json:"text" binding:"required"`
	Date   string `json:"date" binding:"required"`
}

type CreateEventResponseV1 struct {
	EventID string `json:"event_id"`
	Status  string `json:"status"`
}

type UpdateEventRequestV1 struct {
	UserID  int    `json:"user_id" binding:"required"`
	EventID string `json:"event_id" binding:"required"`
	Text    string `json:"text" binding:"required"`
	Date    string `json:"date" binding:"required"`
}

type UpdateEventResponseV1 struct {
	Status string `json:"status"`
}

type DeleteEventRequestV1 struct {
	UserID  int    `json:"user_id" binding:"required"`
	EventID string `json:"event_id" binding:"required"`
}

type DeleteEventResponseV1 struct {
	Status string `json:"status"`
}

type ListEventsRequestV1 struct {
	UserID int    `json:"user_id" binding:"required"`
	Date   string `json:"date" binding:"required"`
}

type EventDtoV1 struct {
	EventID string    `json:"event_id"`
	Text    string    `json:"text"`
	Date    time.Time `json:"date"`
}

type ListEventsResponse struct {
	Events []EventDtoV1 `json:"events"`
}
