package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"events_api/internal/models"
)

func getEventsByDate(events []models.Event, date time.Time) []models.EventResponse {
	var matchingEvents []models.EventResponse = make([]models.EventResponse, 0)

	// Iterate through all events and check if the date matches
	for _, event := range events {

		// Check if the event's date matches the target date
		if event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			matchingEvents = append(matchingEvents, models.EventResponse{
				Category:    event.Category,
				Date:        event.Date.Format("2006-01-02"),
				Description: event.Description,
			})
		}
	}
	return matchingEvents
}

func EventHandler(allEvents []models.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Only Get is allowed for now
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed or implemented", http.StatusMethodNotAllowed)
			return
		}

		if r.URL.Query().Get("date") == "" {
			http.Error(w, "Date parameter is required", http.StatusBadRequest)
			return
		}

		// Parse date parameter
		parsedDate, err := time.Parse("01-02", r.URL.Query().Get("date"))
		if err != nil {
			http.Error(w, "Invalid date format. Use MM-DD", http.StatusBadRequest)
			return
		}

		// Find matching events based on given date
		events := getEventsByDate(allEvents, parsedDate)

		w.Header().Set("Content-Type", "application/json")
		if len(events) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		// Encode the matching event data into json format
		json.NewEncoder(w).Encode(events)

		// DEBUG INFORMATION
		fmt.Printf("Responded with %d events for date %s\n", len(events), parsedDate.Format("01-02"))

	}
}
