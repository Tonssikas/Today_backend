package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"events_api/internal/models"
)

func ReadCSV(filePath string) ([]models.Event, error) {

	//filePath = "./fake10k.csv"

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var events []models.Event

	// Assuming event format is the following
	// 0 = date
	// 1 = description
	// 2 = category
	for _, record := range records {

		date, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			continue // Skip records with invalid date format
		}
		description := record[1]
		category := record[2]

		event := models.Event{
			Category:    category,
			Date:        date,
			Description: description,
		}

		events = append(events, event)

		fmt.Printf("Parsed event: %s, Category: %s, Date: %s\n", description, category, date.Format("2006-01-02"))
	}

	return events, nil
}
