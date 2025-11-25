package models

import "time"

type Data struct {
	Meta  Meta
	Event Event
}

type Meta struct {
	EventID    string
	UserID     int
	EventDate  time.Time
	UpdateDate time.Time
}

type Event struct {
	Text string
}
