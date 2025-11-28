package v1

type CreateRequestV1 struct {
	UserID    int    `json:"user_id" binding:"required"`
	EventDate string `json:"event_date" binding:"required"`
	Text      string `json:"text" binding:"required"`
}

type CreateResponseV1 struct {
	EventID string `json:"event_id"`
}

type UpdateRequestV1 struct {
	UserID      int    `json:"user_id" binding:"required"`
	EventID     string `json:"event_id" binding:"required"`
	CurrentDate string `json:"current_date" binding:"required"`
	NewDate     string `json:"new_date" binding:"required"`
	Text        string `json:"text" binding:"required"`
}

type UpdateResponseV1 struct {
	Updated bool `json:"updated"`
}

type DeleteRequestV1 struct {
	UserID    int    `json:"user_id" binding:"required"`
	EventID   string `json:"event_id" binding:"required"`
	EventDate string `json:"event_date" binding:"required"`
}

type DeleteResponseV1 struct {
	Deleted bool `json:"deleted"`
}

type EventDtoV1 struct {
	Text string `json:"text"`
}

type ListOfEventsResponseV1 struct {
	Events []EventDtoV1 `json:"events"`
}
