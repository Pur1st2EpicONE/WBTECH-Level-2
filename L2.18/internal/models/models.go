package models

import "time"

type Data struct {
	Meta  Meta
	Event Event
}

type Meta struct {
	UserID      int
	EventID     string
	CurrentDate time.Time
	NewDate     time.Time
}

type Event struct {
	Text string
}
