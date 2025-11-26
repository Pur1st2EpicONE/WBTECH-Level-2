package models

import "time"

type Data struct {
	Meta  Meta
	Event Event
}

type Meta struct {
	EventID   string
	UserID    int
	EventDate time.Time
	NewDate   time.Time
}

type Event struct {
	Text string
}
