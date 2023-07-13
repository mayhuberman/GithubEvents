package handlers

import (
	"encoding/json"
	"fmt"
	"githubEventsListener/models"
	"io"
	"log"
	"net/http"
	"sort"
)

var eventTypes = make(map[string]int)

var actorsList models.LinkedList
var actorsMap = make(map[string]int)
var actorsSize = 10

var urlsList models.LinkedList
var urlsSize = 10

var emails = make(map[string]bool)

func ReturnData(w http.ResponseWriter) {
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"eventTypes": eventTypes,
		"actors":     actorsList.ConvertToSlice(),
		"urls":       urlsList.ConvertToSlice(),
		"emails":     getMapKeys(emails),
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln("An error occurred while trying to encode json: ", data, err)
		return
	}
}

func ReturnReposStars(w http.ResponseWriter) {
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	repoStars := getRepoStars(urlsList) // change implementation of getRepoStars

	data := map[string]interface{}{
		"repoStars": repoStars.ConvertUrlStarsToStrings(),
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln("An error occurred while trying to encode json: ", data, err)
		return
	}
}

func ExtractData() {
	events := GetEvents()
	index := 0

	for _, event := range events {

		eventTypes[event.Type] = eventTypes[event.Type] + 1

		var commits = event.Payload.Commits
		for _, commit := range commits {
			if commit.Author.Email != "" {
				emails[commit.Author.Email] = true
			}
		}

		actorName := event.Actor.Login
		if actorsList.Size() < actorsSize {
			if !actorsList.Search(actorName) {
				actorsList.Append(actorName)
				actorsMap[actorName] = index
				index = index + 1
			} else {
				currIndex := actorsMap[actorName]
				actorsMap[actorName] = index - 1
				for key, value := range actorsMap {
					if value >= currIndex && key != actorName {
						actorsMap[key] = value - 1
					}
				}
			}
		} else {
			if !actorsList.Search(actorName) {
				for key, value := range actorsMap {
					if value == 0 {
						actorsList.Detach(key)
						delete(actorsMap, key)
						break
					}
				}
				for key, value := range actorsMap {
					actorsMap[key] = value - 1
				}
				actorsList.Append(actorName)
				actorsMap[actorName] = actorsSize - 1
			} else {
				currIndex := actorsMap[actorName]
				actorsMap[actorName] = actorsSize - 1
				for key, value := range actorsMap {
					if value > currIndex && key != actorName {
						actorsMap[key] = value - 1
					}
				}
			}
		}

		url := event.Repo.URL
		if urlsList.Size() >= urlsSize {
			urlsList.DetachHead()
		}
		urlsList.Append(url)
	}

	println("***************************************************************************")
	actorsList.Print()
	println("***************************************************************************")
	println("events length: ", len(events))
	println("eventTypes: ", eventTypes)
	println("***************************************************************************")
	urlsList.Print()
}

func getRepoStars(urlsList models.LinkedList) models.UrlStarsSlice {
	var urlsStars []models.UrlStars

	for urlsList.Size() > 0 {
		url := urlsList.GetHeadDataAndMoveNext()
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln("An error occurred while trying to get the repo stars from Github. url: ", url, err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("An error occurred while trying to read the response from Github.: ", err)
			return nil
		}
		var repo models.Repo
		err = json.Unmarshal(body, &repo)
		if err != nil {
			fmt.Println("An error occurred while trying to unmarshal the response from Github.: ", err)
			return nil
		}
		newPair := models.UrlStars{Url: url, Stars: repo.StarsCount}
		urlsStars = append(urlsStars, newPair)
	}
	sort.Slice(urlsStars, func(i, j int) bool {
		return urlsStars[i].Stars < urlsStars[j].Stars
	})
	return urlsStars
}
