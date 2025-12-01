package v1

type CreateRequestV1 struct {
	UserID    int    `json:"user_id"`
	EventDate string `json:"date"`
	Text      string `json:"text,omitempty"`
}

type CreateResponseV1 struct {
	EventID string `json:"event_id"`
}

type UpdateRequestV1 struct {
	UserID  int    `json:"user_id"`
	EventID string `json:"event_id"`
	Text    string `json:"text,omitempty"`
	NewDate string `json:"new_date,omitempty"`
}

type UpdateResponseV1 struct {
	Updated bool `json:"event_updated"`
}

type DeleteRequestV1 struct {
	UserID  int    `json:"user_id" binding:"required"`
	EventID string `json:"event_id" binding:"required"`
}

type DeleteResponseV1 struct {
	Deleted bool `json:"event_deleted"`
}

type EventDtoV1 struct {
	Text      string `json:"text"`
	EventDate string `json:"date"`
	EventID   string `json:"event_id"`
}

type ListOfEventsResponseV1 struct {
	Events []EventDtoV1 `json:"events"`
}
