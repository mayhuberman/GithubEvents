package main

import (
	"fmt"
	"githubEventsListener/clients"
	"githubEventsListener/services"
	"log"
	"net/http"
	"time"
)

func main() {
	go runBackgroundJob()

	mux := http.NewServeMux()

	// Register clients for each API endpoint
	mux.HandleFunc("/events", handlerEvents)
	mux.HandleFunc("/repo-stars", handleReposStars)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func runBackgroundJob() {
	for {
		// Perform the background task
		fmt.Println("Running a background job...")
		events := clients.GetEvents()
		services.ExtractData(events)

		// Sleep for 60 seconds
		time.Sleep(60 * time.Second)
	}
}

func handlerEvents(w http.ResponseWriter, r *http.Request) {
	services.ReturnData(w)
}

func handleReposStars(w http.ResponseWriter, r *http.Request) {
	services.ReturnReposStars(w)
}
