package models

import "time"

type Event struct {
	Meta Meta
	Data Data
}

type Meta struct {
	UserID    int
	EventID   string
	EventDate time.Time
	NewDate   time.Time
}

type Data struct {
	Text string
}
