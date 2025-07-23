package main

import (
	"fmt"
	"net/http"

	"events_api/internal/csv"
	"events_api/internal/handlers"
)

func main() {
	// Load events from CSV
	csvFilePath := "data/events.csv"

	events, err := csv.ReadCSV(csvFilePath)
	if err != nil {
		fmt.Printf("Error reading CSV file: %v\n", err)
		return
	}

	fmt.Printf("Loaded %d events from CSV file\n", len(events))

	// Setup routes - pass events to handler
	http.HandleFunc("/api/v1/event/", handlers.EventHandler(events))

	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
