package models

import "time"

type Event struct {
	EventID string
	UserID  int
	Date    time.Time
	Text    string
}
