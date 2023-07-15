package handlers

import (
	"encoding/json"
	"githubEventsListener/models"
	"io"
	"log"
	"net/http"
)

func GetEvents() []models.Event {
	url := "https://api.github.com/events"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("An error occurred while trying to get the events from Github.: ", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer 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")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("An error occurred while trying to get the events from Github.: ", err)
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("An error occurred while trying to get the events from Github.: ", err)
		return nil
	}
	var events []models.Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil
	}
	return events
}
