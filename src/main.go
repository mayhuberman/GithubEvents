package main

import (
	"fmt"
	"githubEventsListener/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	go runBackgroundJob()

	mux := http.NewServeMux()

	// Register handlers for each API endpoint
	mux.HandleFunc("/events", handler)
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
		handlers.ExtractData()

		// Sleep for 60 seconds
		time.Sleep(300 * time.Second)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	handlers.ReturnData(w)
}

func handleReposStars(w http.ResponseWriter, r *http.Request) {
	handlers.ReturnReposStars(w)
}
