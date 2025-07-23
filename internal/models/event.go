package models

import "time"

type DateOnly time.Time

type Event struct {
	Category    string
	Date        time.Time
	Description string
}

type EventResponse struct {
	Category    string `json:"category"`
	Date        string `json:"date"`
	Description string `json:"description"`
}
